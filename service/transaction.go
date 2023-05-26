package service

import (
	"context"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
)

// CreateTransaction implements Service.
func (s *service) CreateTransaction(ctx context.Context, arg db.CreateTransactionParams) error {
	return s.store.CreateTransaction(ctx, arg)
}

// GetTransactionUserID implements Service.
func (s *service) GetTransactionUserID(ctx context.Context, userID int32) ([]db.Transaction, error) {
	return s.store.GetTransactionUserID(ctx, userID)
}

// GetTransactionWalletID implements Service.
func (s *service) GetTransactionWalletID(ctx context.Context, walletID int32) ([]db.Transaction, error) {
	return s.store.GetTransactionWalletID(ctx, walletID)
}

// GetTransactionWalletByidAndUserID implements Service.
func (s *service) GetTransactionWalletByidAndUserID(ctx context.Context, arg db.GetTransactionWalletByidAndUserIDParams) ([]db.Transaction, error) {
	return s.store.GetTransactionWalletByidAndUserID(ctx, arg)
}
