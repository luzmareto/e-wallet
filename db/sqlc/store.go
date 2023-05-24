package db

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	TRX_TRANSFER   = "TRANSFER:\n"
	TRX_TOPUP      = "TOPUP:\n"
	TRX_WITHDRAWAL = "WITHDRAWAL:\n"
	TRX_PAYMENT    = "PAYMENT:\n"
)

type Store interface {
	Querier
	TransferTransactions(ctx context.Context, arg CreateTransferParams) (TransferResult, error)
	TopupTransactions(ctx context.Context, arg CreateTopUpsParams) (TopupResult, error)
	WithdrawalTransactions(ctx context.Context, arg CreateWithdrawalsParams) (WithdrawalResult, error)
	// MerchantPaymentTransactions()
}

type sqlStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &sqlStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *sqlStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queries := New(tx)
	err = fn(queries)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("err : %v, rb err : %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// // MerchantPayment implements Store.
// func (store *sqlStore) MerchantPaymentTransactions() {
// 	panic("unimplemented")
// }

type TopupResult struct {
	Topup  Topup  `json:"topup_details"`
	Wallet Wallet `json:"wallet"`
}

// Topup implements Store.
func (store *sqlStore) TopupTransactions(ctx context.Context, arg CreateTopUpsParams) (TopupResult, error) {
	var result TopupResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Wallet, err = q.AddWalletBalance(ctx, AddWalletBalanceParams{
			ID:      int64(arg.WalletID),
			Balance: arg.Amount,
		})
		if err != nil {
			return err
		}
		result.Topup, err = q.CreateTopUps(ctx, CreateTopUpsParams{
			UserID:      arg.UserID,
			WalletID:    arg.WalletID,
			Amount:      arg.Amount,
			Description: arg.Description,
		})
		if err != nil {
			return err
		}
		trxArg := CreateTransactionParams{
			UserID:      arg.UserID,
			WalletID:    arg.WalletID,
			Amount:      arg.Amount,
			Description: TRX_TOPUP + arg.Description,
		}
		err = q.CreateTransaction(ctx, trxArg)
		return err
	})
	return result, err
}

type CreateTransferParams struct {
	UserID       int32   `json:"user_id"`
	FromWalletID int32   `json:"from_wallet_id"`
	ToWalletID   int32   `json:"to_wallet_id"`
	Amount       float64 `json:"amount"`
	Description  string  `json:"description"`
}

type TransferResult struct {
	Transfer   Transfer `json:"transfer"`
	FromWallet Wallet   `json:"from_wallet"`
	ToWallet   Wallet   `json:"to_wallet"`
}

// Transfer implements Store.
func (store *sqlStore) TransferTransactions(ctx context.Context, arg CreateTransferParams) (TransferResult, error) {
	var result TransferResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.FromWallet, err = q.AddWalletBalance(ctx, AddWalletBalanceParams{
			ID:      int64(arg.FromWalletID),
			Balance: -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToWallet, err = q.AddWalletBalance(ctx, AddWalletBalanceParams{
			ID:      int64(arg.ToWalletID),
			Balance: arg.Amount,
		})
		if err != nil {
			return err
		}

		err = q.CreateTransaction(ctx, CreateTransactionParams{
			UserID:      arg.UserID,
			WalletID:    arg.FromWalletID,
			Amount:      -arg.Amount,
			Description: TRX_TRANSFER + arg.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateTransaction(ctx, CreateTransactionParams{
			UserID:      arg.UserID,
			WalletID:    arg.ToWalletID,
			Amount:      arg.Amount,
			Description: TRX_TRANSFER + arg.Description,
		})
		if err != nil {
			return err
		}

		trfArg := CreateTransfersParams{
			FromWalletID: arg.FromWalletID,
			ToWalletID:   arg.ToWalletID,
			Amount:       arg.Amount,
			Description:  arg.Description,
		}
		result.Transfer, err = q.CreateTransfers(ctx, trfArg)
		return err
	})

	return result, err
}

type WithdrawalResult struct {
	Withdrawal Withdrawal `json:"withdawal_details"`
	Wallet     Wallet     `json:"wallet"`
}

// Withdrawal implements Store.
func (store *sqlStore) WithdrawalTransactions(ctx context.Context, arg CreateWithdrawalsParams) (WithdrawalResult, error) {
	var result WithdrawalResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Wallet, err = q.AddWalletBalance(ctx, AddWalletBalanceParams{
			ID:      int64(arg.WalletID),
			Balance: -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.Withdrawal, err = q.CreateWithdrawals(ctx, CreateWithdrawalsParams{
			UserID:      arg.UserID,
			WalletID:    arg.WalletID,
			Amount:      arg.Amount,
			Description: arg.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateTransaction(ctx, CreateTransactionParams{
			UserID:      arg.UserID,
			WalletID:    arg.WalletID,
			Amount:      -arg.Amount,
			Description: TRX_WITHDRAWAL + arg.Description,
		})
		return err
	})
	return result, err
}
