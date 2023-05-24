package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/service"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
	"github.com/gin-gonic/gin"
)

type MerchantHandler interface {
	CreateMerchant(ctx *gin.Context)
	DeleteMerchant(ctx *gin.Context)
	GetMerchantById(ctx *gin.Context)
	ListMerchant(ctx *gin.Context)
	GetMerchantByUsername(ctx *gin.Context)
	UpdateMerchant(ctx *gin.Context)
}

type merchantHandler struct {
	service service.Service
}

func NewMerchantHandler(service service.Service) MerchantHandler {
	return &merchantHandler{service: service}
}

type UpdatMerchantsRequest struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

func (h *merchantHandler) UpdateMerchant(ctx *gin.Context) {
	var request UpdatMerchantsRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.UpdatMerchantsParams{
		ID:          request.ID,
		Description: request.Description,
		Address:     request.Address,
	}

	data, err := h.service.UpdatMerchants(ctx, arg)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			ctx.JSON(http.StatusNotFound, gin.H{"error": customErr.Msg})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update merchant"})
		return
	}
	ctx.JSON(responseOK("Success", data))
}

func (h *merchantHandler) GetMerchantByUsername(ctx *gin.Context) {
	payload, err := middleware.GetPayload(ctx)
	if err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	data, err := h.service.GetMerchantsByMerchantsName(ctx, payload.Username)
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

type ListMerchantsRequest struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (h *merchantHandler) ListMerchant(ctx *gin.Context) {

	var request ListMerchantsRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.ListMerchantsParams{
		Limit:  request.Limit,
		Offset: request.Offset,
	}

	data, err := h.service.ListMerchants(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve merchants"})
		return
	}

	ctx.JSON(responseOK("Success", data))
}

func (h *merchantHandler) GetMerchantById(ctx *gin.Context) {
	merchantID := ctx.Param("id")

	// Convert the merchant ID to int64
	id, err := strconv.ParseInt(merchantID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid merchant ID"})
		return
	}

	// Call the service method to retrieve the merchant by ID
	merchant, err := h.service.GetMerchantsById(ctx, id)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			ctx.JSON(http.StatusNotFound, gin.H{"error": customErr.Msg})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve merchant"})
		return
	}

	ctx.JSON(http.StatusOK, merchant)
}

func (h *merchantHandler) DeleteMerchant(ctx *gin.Context) {
	// Extract the merchant ID from the request parameters
	merchantID := ctx.Param("id")

	// Convert the merchant ID to int64
	id, err := strconv.ParseInt(merchantID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid merchant ID"})
		return
	}

	// Call the service method to delete the merchant
	err = h.service.DeleteMerchants(ctx, id)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			ctx.JSON(http.StatusNotFound, gin.H{"error": customErr.Msg})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete merchant"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Merchant deleted successfully"})

}

type CreateMerchantsRequest struct {
	MerchantName string `json:"merchant_name" binding:"required"`
	Description  string `json:"description"`
	Website      string `json:"website"`
	Address      string `json:"address"`
}

func (h *merchantHandler) CreateMerchant(ctx *gin.Context) {
	var req CreateMerchantsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create merchant"})
		return
	}

	ctx.JSON(responseOK("Success", data))
}
