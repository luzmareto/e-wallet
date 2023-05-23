package handler

import (
	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/service"
	"github.com/gin-gonic/gin"
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
	UserID      int32   `json:"user_id"`
	WalletID    int32   `json:"wallet_id"`
	Amount      float64 `json:"amount"`
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
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}
	ctx.JSON(responseOK("Success", data))

}
