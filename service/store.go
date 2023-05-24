package service

import (
	"context"
	"database/sql"
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
	wallet, err := s.queries.GetWalletById(ctx, int64(arg.WalletID))
	if err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("wallet with id %d not found", arg.WalletID),
			Err: err,
		}
		return db.WithdrawalResult{}, cstErr
	}

	user, err := s.queries.GetUserById(ctx, int64(arg.UserID))
	if err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("user with id %d not found", arg.WalletID),
			Err: err,
		}
		return db.WithdrawalResult{}, cstErr
	}

	if wallet.UserID != int32(user.ID) {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("wallet with id %d does not belong to user with id %d", wallet.ID, user.ID),
			Err: fmt.Errorf("unauthorized"),
		}
		return db.WithdrawalResult{}, cstErr
	}

	return s.store.WithdrawalTransactions(ctx, arg)
}

// TransferTransactions implements Service.
func (s *service) TransferTransactions(ctx context.Context, arg db.CreateTransferParams) (db.TransferResult, error) {
	var result db.TransferResult
	var err error

	walletFrom, err := s.queries.GetWalletById(ctx, int64(arg.FromWalletID))
	if err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("wallet with id %d not found", arg.FromWalletID),
			Err: err,
		}
		return db.TransferResult{}, cstErr
	}

	walletTo, err := s.queries.GetWalletById(ctx, int64(arg.ToWalletID))
	if err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("wallet with id %d not found", arg.ToWalletID),
			Err: err,
		}
		return db.TransferResult{}, cstErr
	}

	if walletFrom.Currency != walletTo.Currency {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("cannt transfer from %s wallet to %s wallet", walletFrom.Currency, walletTo.Currency),
			Err: err,
		}
		return db.TransferResult{}, cstErr
	}

	if walletFrom.Balance < arg.Amount {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("wallet %d balance not enought %s. %.2f", arg.FromWalletID, walletFrom.Currency, walletFrom.Balance),
			Err: sql.ErrConnDone,
		}
		return db.TransferResult{}, cstErr
	}

	return result, err
}
