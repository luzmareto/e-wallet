package service

import (
	"context"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
)

// CreateTopUps implements Service
func (s *service) CreateTopUps(ctx context.Context, arg db.CreateTopUpsParams) (db.Topup, error) {
	return s.store.CreateTopUps(ctx, arg)
}
