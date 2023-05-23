package handler

import (
	"github.com/gin-gonic/gin"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/service"
)

type WithdrawalHandler interface {
	CreateWithdrawal(ctx *gin.Context)
}

type withdrawalHandler struct {
	service service.Service
}

func NewWithdrawalHandler(service service.Service) WithdrawalHandler {
	return &withdrawalHandler{service: service}
}

type createwithdrawalsRequest struct {
	UserID      int32   `json:"user_id" binding:"required"`
	WalletID    int32   `json:"wallet_id" binding:"required,min=1"`
	Amount      float64 `json:"amount" binding:"required,min=10000"`
	Description string  `json:"description"`
}

// Register implements UserHandler.
func (h *withdrawalHandler) CreateWithdrawal(ctx *gin.Context) {
	var req createwithdrawalsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	arg := db.CreatewithdrawalsParams{
		UserID:      req.UserID,
		WalletID:    req.WalletID,
		Amount:      req.Amount,
		Description: req.Description,
	}

	data, err := h.service.Createwithdrawals(ctx, arg)
	if err != nil {
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}
	ctx.JSON(responseOK("Success", data))
}
