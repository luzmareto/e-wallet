package service

import (
	"context"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
)

// Topup implements Service.
func (s *service) TopupTransactions(ctx context.Context, arg db.CreateTopUpsParams) (db.TopupResult, error) {
	return s.store.TopupTransactions(ctx, arg)
}

// WithdrawalTransactions implements Service.
func (s *service) WithdrawalTransactions(ctx context.Context, arg db.CreateWithdrawalsParams) (db.WithdrawalResult, error) {
	return s.store.WithdrawalTransactions(ctx, arg)
}
