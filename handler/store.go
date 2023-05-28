package handler

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/service"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

type StoreHandler interface {
	TopupTransactions(ctx *gin.Context)
	WithdrawalTransactions(ctx *gin.Context)
	TransferTransactions(ctx *gin.Context)
	MerchantPaymentTransactions(ctx *gin.Context)
	WalletHistory(ctx *gin.Context)
}

type storeHandler struct {
	service service.Service
}

func NewStoreHandler(service service.Service) StoreHandler {
	return &storeHandler{service: service}
}

type topupTransactionsRequest struct {
	UserID      int32   `json:"-"`
	WalletID    int32   `json:"wallet_id" binding:"required,min=1"`
	Amount      float64 `json:"amount" binding:"min=10000,max=10000000"`
	Description string  `json:"description"`
}

// TopupTransactions implements StoreHandler.
func (h *storeHandler) TopupTransactions(ctx *gin.Context) {
	payload, err := middleware.GetPayload(ctx)
	if err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	user, err := h.service.GetUserByUserName(ctx, payload.Username)
	newErr := utils.CastError(err)

	if err != nil {
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(newErr.Msg))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	var req topupTransactionsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}

	arg := db.CreateTopUpsParams{
		UserID:      int32(user.ID),
		WalletID:    req.WalletID,
		Amount:      req.Amount,
		Description: req.Description,
	}
	data, err := h.service.TopupTransactions(ctx, arg)
	newErr = utils.CastError(err)

	if err != nil {
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(newErr.Msg))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}
	ctx.JSON(responseOK("Success", data))
}

type withdrawalsRequest struct {
	UserID      int32   `json:"-"`
	WalletID    int32   `json:"wallet_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

// WithdrawalTransactions implements StoreHandler.
func (h *storeHandler) WithdrawalTransactions(ctx *gin.Context) {
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
	var req withdrawalsRequest
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
	data, err := h.service.WithdrawalTransactions(ctx, arg)
	newErr = utils.CastError(err)

	if err != nil {
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(newErr.Msg))
			return
		}
		if newErr.Err.Error() == "unauthorized" {
			ctx.JSON(responseUnauthorized(newErr.Msg))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}
	ctx.JSON(responseOK("Success", data))
}

type createTransferTransactionsRequest struct {
	FromWalletID int32   `json:"from_wallet_id" binding:"required,min=1"`
	ToWalletID   int32   `json:"to_wallet_id" binding:"required,min=1"`
	Amount       float64 `json:"amount" binding:"required,min=10000"`
	Description  string  `json:"description"`
}

// TransferTransactions implements StoreHandler.
func (h *storeHandler) TransferTransactions(ctx *gin.Context) {
	payload, err := middleware.GetPayload(ctx)
	if err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	user, err := h.service.GetUserByUserName(ctx, payload.Username)
	if err != nil {
		newErr := utils.CastError(err)
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(newErr.Msg))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	var req createTransferTransactionsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	arg := db.CreateTransferParams{
		UserID:       int32(user.ID),
		FromWalletID: req.FromWalletID,
		ToWalletID:   req.ToWalletID,
		Amount:       req.Amount,
		Description:  req.Description,
	}

	data, err := h.service.TransferTransactions(ctx, arg)
	if err != nil {
		newErr := err.(*utils.CustomError)
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(newErr.Msg))
			return
		}

		if newErr.Err == sql.ErrConnDone {
			ctx.JSON(responseBadRequest(newErr.Msg))
			return
		}

		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	ctx.JSON(responseOK("success", data))
}

type merchantPaymentRequest struct {
	UserID      int32   `json:"-"`
	WalletID    int32   `json:"wallet_id"`
	MerchantID  int64   `json:"merchant_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

// MerchantPaymentTransactions implements StoreHandler.
func (h *storeHandler) MerchantPaymentTransactions(ctx *gin.Context) {
	payload, err := middleware.GetPayload(ctx)
	if err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	user, err := h.service.GetUserByUserName(ctx, payload.Username)
	if err != nil {
		newErr := utils.CastError(err)
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(newErr.Msg))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	var req merchantPaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}

	arg := db.CreateTransactionParams{
		UserID:      int32(user.ID),
		WalletID:    req.WalletID,
		Amount:      req.Amount,
		Description: req.Description,
	}

	err = h.service.MerchantPaymentTransactions(ctx, arg, req.MerchantID)
	if err != nil {
		newErr := err.(*utils.CustomError)
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(newErr.Msg))
			return
		}

		if newErr.Err == sql.ErrConnDone {
			ctx.JSON(responseBadRequest(newErr.Msg))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	ctx.JSON(responseOK("success"))
}

type walletHistoryRequest struct {
	WalletID int32  `form:"wallet_id" binding:"required,min=1"`
	Type     string `form:"type" binding:"required,oneof='transfers' 'transactions'"`
}

// WalletHistory implements StoreHandler.
func (h *storeHandler) WalletHistory(ctx *gin.Context) {
	payload, err := middleware.GetPayload(ctx)
	if err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}
	user, err := h.service.GetUserByUserName(ctx, payload.Username)
	if err != nil {
		newErr := utils.CastError(err)
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(newErr.Msg))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	var req walletHistoryRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(responseBadRequest(err.Error()))
		return
	}

	pathcsv, err := h.service.WalletHistory(ctx, db.GetTransactionWalletByidAndUserIDParams{
		WalletID: req.WalletID,
		UserID:   int32(user.ID),
	}, req.Type)

	defer os.RemoveAll(pathcsv)

	if err != nil {
		newErr := err.(*utils.CustomError)
		if newErr.Err == sql.ErrNoRows {
			ctx.JSON(responseNotFound(newErr.Msg))
			return
		}

		if newErr.Err == sql.ErrConnDone {
			ctx.JSON(responseBadRequest(newErr.Msg))
			return
		}
		ctx.JSON(responseInternalServerError(err.Error()))
		return
	}

	// Set the necessary headers for file download
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.csv", req.Type))
	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Cache-Control", "no-cache")

	fmt.Println(pathcsv)
	ctx.File(pathcsv)

	// ctx.JSON(responseOK("Success", pathcsv))

}
