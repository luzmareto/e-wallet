// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: transation.sql

package db

import (
	"context"
)

const createTransaction = `-- name: CreateTransaction :exec
INSERT INTO transactions (
    user_id,
    wallet_id,
    amount, 
    description,
    transaction_type
) VALUES (
    $1, $2, $3, $4, $5
)
`

type CreateTransactionParams struct {
	UserID          int32   `json:"user_id"`
	WalletID        int32   `json:"wallet_id"`
	Amount          float64 `json:"amount"`
	Description     string  `json:"description"`
	TransactionType string  `json:"transaction_type"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) error {
	_, err := q.db.ExecContext(ctx, createTransaction,
		arg.UserID,
		arg.WalletID,
		arg.Amount,
		arg.Description,
		arg.TransactionType,
	)
	return err
}

const getTransactionUserID = `-- name: GetTransactionUserID :many
SELECT id, user_id, wallet_id, amount, transaction_date, transaction_type, description FROM transactions WHERE user_id = $1
`

func (q *Queries) GetTransactionUserID(ctx context.Context, userID int32) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getTransactionUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.WalletID,
			&i.Amount,
			&i.TransactionDate,
			&i.TransactionType,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTransactionWalletID = `-- name: GetTransactionWalletID :many
SELECT id, user_id, wallet_id, amount, transaction_date, transaction_type, description FROM transactions WHERE wallet_id = $1
`

func (q *Queries) GetTransactionWalletID(ctx context.Context, walletID int32) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getTransactionWalletID, walletID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.WalletID,
			&i.Amount,
			&i.TransactionDate,
			&i.TransactionType,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
