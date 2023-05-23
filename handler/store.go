package handler

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/service"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

type StoreHandler interface {
	TopupTransactions(ctx *gin.Context)
	WithdrawalTransactions(ctx *gin.Context)
}

type storeHandler struct {
	service service.Service
}

func NewStoreHandler(service service.Service) StoreHandler {
	return &storeHandler{service: service}
}

type topupTransactionsRequest struct {
	UserID      int32   `json:"user_id" binding:"required"`
	WalletID    int32   `json:"wallet_id" binding:"required,min=1"`
	Amount      float64 `json:"amount" binding:"min=10000,max=10000000"`
	Description string  `json:"description"`
}

// TopupTransactions implements StoreHandler.
func (h *storeHandler) TopupTransactions(ctx *gin.Context) {
	var req topupTransactionsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}

	arg := db.CreateTopUpsParams{
		UserID:      req.UserID,
		WalletID:    req.WalletID,
		Amount:      req.Amount,
		Description: req.Description,
	}
	data, err := h.service.TopupTransactions(ctx, arg)
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

type withdrawalsRequest struct {
	UserID      int32   `json:"user_id"`
	WalletID    int32   `json:"wallet_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

// WithdrawalTransactions implements StoreHandler.
func (h *storeHandler) WithdrawalTransactions(ctx *gin.Context) {
	var req withdrawalsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	arg := db.CreateWithdrawalsParams{
		UserID:      req.UserID,
		WalletID:    req.WalletID,
		Amount:      req.Amount,
		Description: req.Description,
	}
	data, err := h.service.WithdrawalTransactions(ctx, arg)
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
