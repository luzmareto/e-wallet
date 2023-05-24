package handler

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/service"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

type MerchantHandler interface {
	Register(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetById(ctx *gin.Context)
	GetByName(ctx *gin.Context)
	List(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type merchantHandler struct {
	service service.Service
}

func NewMerchantHandler(service service.Service) MerchantHandler {
	return &merchantHandler{service: service}
}

type updatMerchantRequest struct {
	ID          int64  `json:"id" binding:"required,min=1"`
	Description string `json:"description" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

func (h *merchantHandler) Update(ctx *gin.Context) {
	payload, err := middleware.GetPayload(ctx)
	if err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	if payload.Role != ROLE_ADMIN {
		ctx.JSON(responseUnauthorized("role unauthorized"))
		return
	}
	var req updatMerchantRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}

	arg := db.UpdatMerchantsParams{
		ID:          req.ID,
		Description: req.Description,
		Address:     req.Address,
	}

	data, err := h.service.UpdatMerchants(ctx, arg)
	if err != nil {
		newErr := utils.CastError(err)
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(newErr.Error()))
			return
		}
		ctx.JSON(responseInternalServerError("failed to update merchant"))
		return
	}
	ctx.JSON(responseOK("Success", data))
}

type getMerchantByNameRequest struct {
	Name string `form:"name" binding:"required"`
}

func (h *merchantHandler) GetByName(ctx *gin.Context) {
	var req getMerchantByNameRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	data, err := h.service.GetMerchantsByMerchantsName(ctx, req.Name)
	if err != nil {
		newErr := utils.CastError(err)
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(err.Error()))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}
	ctx.JSON(responseOK("Success", data))
}

type listMerchantsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=100"`
}

func (h *merchantHandler) List(ctx *gin.Context) {
	var req listMerchantsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}

	arg := db.ListMerchantsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	data, err := h.service.ListMerchants(ctx, arg)
	if err != nil {
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	ctx.JSON(responseOK("Success", data))
}

type getMerchantByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (h *merchantHandler) GetById(ctx *gin.Context) {
	var req getMerchantByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	data, err := h.service.GetMerchantsById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(fmt.Sprintf("merchant with id %d not found", req.ID)))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	ctx.JSON(responseOK("Success", data))
}

func (h *merchantHandler) Delete(ctx *gin.Context) {
	payload, err := middleware.GetPayload(ctx)
	if err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	if payload.Role != ROLE_ADMIN {
		ctx.JSON(responseUnauthorized("role unauthorized"))
		return
	}

	var req getMerchantByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}

	err = h.service.DeleteMerchants(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(fmt.Sprintf("merchant with id %d not found", req.ID)))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	ctx.JSON(responseOK("Success", nil))

}

type createMerchantRequest struct {
	MerchantName string `json:"merchant_name" binding:"required"`
	Description  string `json:"description"`
	Website      string `json:"website"`
	Address      string `json:"address"`
}

func (h *merchantHandler) Register(ctx *gin.Context) {
	payload, err := middleware.GetPayload(ctx)
	if err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	if payload.Role != ROLE_ADMIN {
		ctx.JSON(responseUnauthorized("role unauthorized"))
		return
	}

	var req createMerchantRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}

	arg := db.CreateMerchantsParams{
		MerchantName: req.MerchantName,
		Description:  req.Description,
		Website:      req.Website,
		Address:      req.Address,
	}

	data, err := h.service.CreateMerchants(ctx, arg)
	if err != nil {
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	ctx.JSON(responseOK("Success", data))
}
