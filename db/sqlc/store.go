package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	// Transfer()
	TopupTransactions(ctx context.Context, arg CreateTopUpsParams) (TopupResult, error)
	WithdrawalTransactions(ctx context.Context, arg CreateWithdrawalsParams) (WithdrawalResult, error)
	// MerchantPayment()
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
// func (store *sqlStore) MerchantPayment() {
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

		err = q.CreateTransaction(ctx, CreateTransactionParams{
			UserID:   arg.UserID,
			WalletID: arg.WalletID,
			Amount:   arg.Amount,
			Description: sql.NullString{
				String: arg.Description,
			},
		})
		return err
	})
	return result, err
}

// // Transfer implements Store.
// func (store *sqlStore) Transfer() {
// 	panic("unimplemented")
// }

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
			UserID:   arg.UserID,
			WalletID: arg.WalletID,
			Amount:   arg.Amount,
			Description: sql.NullString{
				String: arg.Description,
			},
		})
		return err
	})
	return result, err
}
