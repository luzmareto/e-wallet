package handler

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/service"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

type WalletHandler interface {
	AddWalletBalance(ctx *gin.Context)
	CreateWallets(ctx *gin.Context)
}

type walletHandler struct {
	service service.Service
}

func NewWalletHandler(service service.Service) WalletHandler {
	return &walletHandler{service: service}
}

type walletIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type addWalletBalanceRequest struct {
	Balance float64 `json:"balance"`
}

// AddWalletBalance implements WalletHandler
func (h *walletHandler) AddWalletBalance(ctx *gin.Context) {
	var reqId walletIdRequest
	if err := ctx.ShouldBindUri(&reqId); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}

	var req addWalletBalanceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}

	arg := db.AddWalletBalanceParams{
		ID:      reqId.ID,
		Balance: req.Balance,
	}

	data, err := h.service.AddWalletBalance(ctx, arg)
	if err != nil {
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}
	ctx.JSON(responseOK("Success", data))
}

type createWalletsRequest struct {
	Balance  float64 `json:"balance,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

// CreateWallets implements WalletHandler
func (h *walletHandler) CreateWallets(ctx *gin.Context) {
	var req createWalletsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
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

	arg := db.CreateWalletsParams{
		UserID:   int32(user.ID),
		Balance:  req.Balance,
		Currency: req.Currency,
	}

	data, err := h.service.CreateWallets(ctx, arg)
	newErr = utils.CastError(err)

	if err != nil {
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(err.Error()))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}
	ctx.JSON(responseCreated("Success Created", data))

}
