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
	if _, err := s.store.GetWalletById(ctx, int64(arg.WalletID)); err != nil {
		if err == sql.ErrNoRows {
			cstErr := &utils.CustomError{
				Msg: fmt.Sprintf("wallet with id %d not found", arg.WalletID),
				Err: err,
			}
			return db.TopupResult{}, cstErr
		}
		return db.TopupResult{}, err
	}
	return s.store.TopupTransactions(ctx, arg)
}

// WithdrawalTransactions implements Service.
func (s *service) WithdrawalTransactions(ctx context.Context, arg db.CreateWithdrawalsParams) (db.WithdrawalResult, error) {
	wallet, err := s.store.GetWalletById(ctx, int64(arg.WalletID))
	if err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("wallet with id %d not found", arg.WalletID),
			Err: err,
		}
		return db.WithdrawalResult{}, cstErr
	}

	user, err := s.store.GetUserById(ctx, int64(arg.UserID))
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

	walletFrom, err := s.store.GetWalletById(ctx, int64(arg.FromWalletID))
	if err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("wallet with id %d not found", arg.FromWalletID),
			Err: err,
		}
		return db.TransferResult{}, cstErr
	}

	if walletFrom.UserID != arg.UserID {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("wallet with id %d not yours", arg.FromWalletID),
			Err: sql.ErrConnDone,
		}
		return db.TransferResult{}, cstErr
	}

	walletTo, err := s.store.GetWalletById(ctx, int64(arg.ToWalletID))
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
			Err: sql.ErrConnDone,
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

	result, err = s.store.TransferTransactions(ctx, arg)
	if err != nil {
		return db.TransferResult{}, err
	}

	return result, err
}

// MerchantPaymentTransactions implements Service.
func (s *service) MerchantPaymentTransactions(ctx context.Context, arg db.CreateTransactionParams, merchantID int64) error {
	wallet, err := s.store.GetWalletById(ctx, int64(arg.WalletID))
	if err != nil {
		if err == sql.ErrNoRows {
			cstErr := &utils.CustomError{
				Msg: fmt.Sprintf("wallet with id %d not found", arg.WalletID),
				Err: err,
			}
			return cstErr
		}
		return err
	}
	if wallet.Balance < arg.Amount {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("wallet %d balance not enought %s. %.2f", arg.WalletID, wallet.Currency, wallet.Balance),
			Err: sql.ErrConnDone,
		}
		return cstErr
	}
	return s.store.MerchantPaymentTransactions(ctx, arg, merchantID)
}

// WalletHistoryGenerateCSV implements Service.
func (s *service) WalletHistoryGenerateCSV(ctx context.Context, arg db.GetTransactionWalletByidAndUserIDParams) (db.WalletHistoryResult, error) {
	return s.store.WalletHistoryGenerateCSV(ctx, arg)
}
