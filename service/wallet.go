package service

import (
	"context"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
)

const (
	CurrencyUSD = "USD"
	CurrencyIDR = "IDR"
	CurrencySGD = "SGD"
)

// CreateWallets implements Service
func (s *service) CreateWallets(ctx context.Context, arg db.CreateWalletsParams) (db.Wallet, error) {
	if arg.Currency == "" {
		arg.Currency = CurrencyIDR
	}
	return s.queries.CreateWallets(ctx, arg)
}

// AddWalletBalance implements Service
func (s *service) AddWalletBalance(ctx context.Context, arg db.AddWalletBalanceParams) (db.Wallet, error) {
	return s.queries.AddWalletBalance(ctx, arg)
}
