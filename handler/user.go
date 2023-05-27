package handler

import (
	"database/sql"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/service"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	GetByUsername(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	List(ctx *gin.Context)
	Update(ctx *gin.Context)
	UploadIDCard(ctx *gin.Context)
}

type userHandler struct {
	service   service.Service
	config    utils.Config
	awsClient utils.AWSS3Client
}

func NewUserHandler(service service.Service, config utils.Config) UserHandler {
	awsClient := utils.NewAWSS3Client(utils.Config{
		AwsAccessKeyID: config.AwsAccessKeyID,
		AwsSecretKeyID: config.AwsSecretKeyID,
		AwsRegion:      config.AwsRegion,
		AwsS3Bucket:    config.AwsS3Bucket,
	})

	return &userHandler{
		service:   service,
		config:    config,
		awsClient: awsClient,
	}
}

type listUsersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=100"`
}

// List implements UserHandler.
func (h *userHandler) List(ctx *gin.Context) {
	payload, err := middleware.GetPayload(ctx)
	if err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	if payload.Role != ROLE_ADMIN {
		ctx.JSON(responseUnauthorized("role unauthorized"))
		return
	}
	var req listUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	data, err := h.service.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}
	ctx.JSON(responseOK("Success", data))
}

type getByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// GetByID implements UserHandler.
func (h *userHandler) GetByID(ctx *gin.Context) {
	payload, err := middleware.GetPayload(ctx)
	if err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	if payload.Role != ROLE_ADMIN {
		ctx.JSON(responseUnauthorized("role unauthorized"))
		return
	}
	var req getByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	data, err := h.service.GetUserById(ctx, req.ID)
	newErr := utils.CastError(err)

	if err != nil {
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(err.Error()))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}
	ctx.JSON(responseOK("Success", data))
}

// GetByUsername implements UserHandler.
func (h *userHandler) GetByUsername(ctx *gin.Context) {
	payload, err := middleware.GetPayload(ctx)
	if err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	data, err := h.service.GetUserByUserName(ctx, payload.Username)
	newErr := utils.CastError(err)

	if err != nil {
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(err.Error()))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}
	ctx.JSON(responseOK("Success", data))
}

type createUsersRequest struct {
	Username    string `json:"username" binding:"required,min=3"`
	Password    string `json:"password" binding:"required,min=8"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required,numeric,min=8,max=13"`
}

// Register implements UserHandler.
func (h *userHandler) Register(ctx *gin.Context) {
	var req createUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	arg := db.CreateUsersParams{
		Username:    req.Username,
		Password:    req.Password,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
	data, err := h.service.CreateUsers(ctx, arg)
	if err != nil {
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}
	ctx.JSON(responseCreated("Success Created", data))
}

type updateUserRequest struct {
	ID          int64  `json:"-"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required,numeric,min=8,max=13"`
}

// Update implements UserHandler.
func (h *userHandler) Update(ctx *gin.Context) {
	payload, err := middleware.GetPayload(ctx)
	if err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	user, err := h.service.GetUserByUserName(ctx, payload.Username)
	if err != nil {
		newErr := utils.CastError(err)
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(err.Error()))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}

	arg := db.UpdateUsersParams{
		ID:          user.ID,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
	data, err := h.service.UpdateUsers(ctx, arg)
	if err != nil {
		newErr := utils.CastError(err)
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(err.Error()))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}
	ctx.JSON(responseOK("Success Update", data))
}

type uploadIDRequest struct {
	File *multipart.FileHeader `form:"file"`
}

// UploadIDCard implements UserHandler.
func (h *userHandler) UploadIDCard(ctx *gin.Context) {
	// Retrieve the file from the request
	var req uploadIDRequest
	if err := ctx.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}

	filename := utils.RandomFileName(req.File)

	// Save the file to the server
	err := ctx.SaveUploadedFile(req.File, utils.DIRECTORY_UPLOADS+"/"+filename)
	if err != nil {
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	uploadedFile, err := h.awsClient.Upload(ctx, req.File, "id-cards")
	if err != nil {
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	ctx.JSON(responseOK("Success", gin.H{
		"uploaded_file": uploadedFile,
	}))
}
