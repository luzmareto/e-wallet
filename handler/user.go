package handler

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/service"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	GetByUsername(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	List(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type userHandler struct {
	service service.Service
}

func NewUserHandler(service service.Service) UserHandler {
	return &userHandler{service: service}
}

type listUsersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// List implements UserHandler.
func (h *userHandler) List(ctx *gin.Context) {
	var req listUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	fmt.Println("ARG LIST USER: ", arg)
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
	var req getByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	data, err := h.service.GetUserById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(fmt.Sprintf("User with ID %d not found", req.ID)))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}
	ctx.JSON(responseOK("Success", data))
}

type getUserByUsernameRequest struct {
	Username string `uri:"username" binding:"required"`
}

// GetByUsername implements UserHandler.
func (h *userHandler) GetByUsername(ctx *gin.Context) {
	var req getUserByUsernameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	data, err := h.service.GetUserByUserName(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(fmt.Sprintf("User with Username %s not found", req.Username)))
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
	ID          int64  `json:"id" binding:"required,min=1"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required,numeric,min=8,max=13"`
}

// Update implements UserHandler.
func (h *userHandler) Update(ctx *gin.Context) {

	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}

	arg := db.UpdateUsersParams{
		ID:          req.ID,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
	data, err := h.service.UpdateUsers(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(fmt.Sprintf("User with ID %d not found", req.ID)))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}
	ctx.JSON(responseOK("Success Update", data))
}
