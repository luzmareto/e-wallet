// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddMerchantBalance(ctx context.Context, arg AddMerchantBalanceParams) (Merchant, error)
	AddWalletBalance(ctx context.Context, arg AddWalletBalanceParams) (Wallet, error)
	CreateMerchants(ctx context.Context, arg CreateMerchantsParams) (Merchant, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateTopUps(ctx context.Context, arg CreateTopUpsParams) (Topup, error)
	CreateTransaction(ctx context.Context, arg CreateTransactionParams) error
	CreateTransfers(ctx context.Context, arg CreateTransfersParams) (Transfer, error)
	CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error)
	CreateWallets(ctx context.Context, arg CreateWalletsParams) (Wallet, error)
	CreateWithdrawals(ctx context.Context, arg CreateWithdrawalsParams) (Withdrawal, error)
	DeleteMerchants(ctx context.Context, id int64) error
	DeleteUsers(ctx context.Context, id int64) error
	GetMerchantsById(ctx context.Context, id int64) (Merchant, error)
	GetMerchantsByMerchantsName(ctx context.Context, merchantName string) (Merchant, error)
	GetSessions(ctx context.Context, id uuid.UUID) (Session, error)
	GetTransactionUserID(ctx context.Context, userID int32) ([]Transaction, error)
	GetTransactionWalletID(ctx context.Context, walletID int32) ([]Transaction, error)
	GetTransfersByFromWalletID(ctx context.Context, fromWalletID int32) ([]Transfer, error)
	GetTransfersByFromWalletIdAndToWalletId(ctx context.Context, arg GetTransfersByFromWalletIdAndToWalletIdParams) ([]Transfer, error)
	GetUserById(ctx context.Context, id int64) (User, error)
	GetUserByUserName(ctx context.Context, username string) (User, error)
	GetWalletById(ctx context.Context, id int64) (Wallet, error)
	ListMerchants(ctx context.Context, arg ListMerchantsParams) ([]Merchant, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdatMerchants(ctx context.Context, arg UpdatMerchantsParams) (Merchant, error)
	UpdateUsers(ctx context.Context, arg UpdateUsersParams) (User, error)
	UpdateUsersPassword(ctx context.Context, arg UpdateUsersPasswordParams) error
}

var _ Querier = (*Queries)(nil)
