package service

import (
	"context"
	"fmt"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

// Topup implements Service.
func (s *service) TopupTransactions(ctx context.Context, arg db.CreateTopUpsParams) (db.TopupResult, error) {
	if _, err := s.queries.GetWalletById(ctx, int64(arg.WalletID)); err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("wallet with id %d not found", arg.WalletID),
			Err: err,
		}
		return db.TopupResult{}, cstErr
	}
	return s.store.TopupTransactions(ctx, arg)
}

// WithdrawalTransactions implements Service.
func (s *service) WithdrawalTransactions(ctx context.Context, arg db.CreateWithdrawalsParams) (db.WithdrawalResult, error) {
	if _, err := s.queries.GetWalletById(ctx, int64(arg.WalletID)); err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("wallet with id %d not found", arg.WalletID),
			Err: err,
		}
		return db.WithdrawalResult{}, cstErr
	}
	return s.store.WithdrawalTransactions(ctx, arg)
}
