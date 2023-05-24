package service

import (
	"context"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
)

// CreateTransfers implements Service.
func (s *service) CreateTransfers(ctx context.Context, arg db.CreateTransfersParams) (db.Transfer, error) {
	return s.queries.CreateTransfers(ctx, arg)
}

// GetTransfersByFromWalletID implements Service.
func (s *service) GetTransfersByFromWalletID(ctx context.Context, fromWalletID int32) ([]db.Transfer, error) {
	return s.queries.GetTransfersByFromWalletID(ctx, fromWalletID)
}

// GetTransfersByFromWalletIdAndToWalletId implements Service.
func (s *service) GetTransfersByFromWalletIdAndToWalletId(ctx context.Context, arg db.GetTransfersByFromWalletIdAndToWalletIdParams) ([]db.Transfer, error) {
	return s.queries.GetTransfersByFromWalletIdAndToWalletId(ctx, arg)
}
