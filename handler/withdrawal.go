package handler

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/service"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
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
	WalletID    int32   `json:"wallet_id" binding:"required,min=1"`
	Amount      float64 `json:"amount" binding:"required,min=10000"`
	Description string  `json:"description"`
}

// Register implements UserHandler.
func (h *withdrawalHandler) CreateWithdrawal(ctx *gin.Context) {
	payload, err := middleware.GetPayload(ctx)
	if err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	user, err := h.service.GetUserByUserName(ctx, payload.Username)
	newErr := utils.CastError(err)
	if err != nil {
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(err.Error()))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	var req createwithdrawalsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	arg := db.CreateWithdrawalsParams{
		UserID:      int32(user.ID),
		WalletID:    req.WalletID,
		Amount:      req.Amount,
		Description: req.Description,
	}

	data, err := h.service.CreateWithdrawals(ctx, arg)
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
