package service

import (
	"context"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
)

// Createwithdrawals implements Service
func (s *service) CreateWithdrawals(ctx context.Context, arg db.CreateWithdrawalsParams) (db.Withdrawal, error) {
	return s.queries.CreateWithdrawals(ctx, arg)
}
