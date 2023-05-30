package service

import (
	"context"
	"fmt"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

const (
	CurrencyUSD = "USD"
	CurrencyIDR = "IDR"
	CurrencySGD = "SGD"
)

// CreateWallets implements Service
func (s *service) CreateWallets(ctx context.Context, arg db.CreateWalletsParams) (db.Wallet, error) {
	if _, err := s.store.GetUserById(ctx, int64(arg.UserID)); err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("user with id %d not found", arg.UserID),
			Err: err,
		}
		return db.Wallet{}, cstErr
	}
	if arg.Currency == "" {
		arg.Currency = CurrencyIDR
	}
	return s.store.CreateWallets(ctx, arg)
}

// GetWalletById implements Service.
func (s *service) GetWalletById(ctx context.Context, id int64) (db.Wallet, error) {
	wallet, err := s.store.GetWalletById(ctx, id)
	if err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("wallet with id %d not found", id),
			Err: err,
		}
		return db.Wallet{}, cstErr
	}
	return wallet, nil
}

// AddWalletBalance implements Service
func (s *service) AddWalletBalance(ctx context.Context, arg db.AddWalletBalanceParams) (db.Wallet, error) {
	return s.store.AddWalletBalance(ctx, arg)
}
