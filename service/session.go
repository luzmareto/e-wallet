package service

import (
	"context"

	"github.com/google/uuid"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
)

// CreateSession implements Service.
func (s *service) CreateSession(ctx context.Context, arg db.CreateSessionParams) (db.Session, error) {
	return s.queries.CreateSession(ctx, arg)
}

// GetSessions implements Service.
func (s *service) GetSessions(ctx context.Context, id uuid.UUID) (db.Session, error) {
	return s.queries.GetSessions(ctx, id)
}
