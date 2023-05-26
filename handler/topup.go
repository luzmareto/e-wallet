package handler

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/service"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

type TopUpHandler interface {
	CreateTopUp(ctx *gin.Context)
}

type topUpHandler struct {
	service service.Service
}

func NewTopUpHandler(service service.Service) TopUpHandler {
	return &topUpHandler{service: service}
}

type CreateTopUpsRequest struct {
	UserID      int32   `json:"user_id" binding:"required"`
	WalletID    int32   `json:"wallet_id" binding:"required,min=1"`
	Amount      float64 `json:"amount" binding:"min=10000,max=10000000"`
	Description string  `json:"description"`
}

func (h *topUpHandler) CreateTopUp(ctx *gin.Context) {
	var req CreateTopUpsRequest
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

	data, err := h.service.CreateTopUps(ctx, arg)
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
