package service

import (
	"context"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

// CreateUsers implements Service.
func (s *service) CreateUsers(ctx context.Context, arg db.CreateUsersParams) (db.User, error) {
	hashedPassword, _ := utils.HashPassword(arg.Password)
	arg.Password = hashedPassword
	return s.queries.CreateUsers(ctx, arg)
}

// DeleteUsers implements Service.
func (s *service) DeleteUsers(ctx context.Context, id int64) error {
	return s.queries.DeleteUsers(ctx, id)
}

// GetUserById implements Service.
func (s *service) GetUserById(ctx context.Context, id int64) (db.User, error) {
	return s.queries.GetUserById(ctx, id)
}

// GetUserByUserName implements Service.
func (s *service) GetUserByUserName(ctx context.Context, username string) (db.User, error) {
	return s.queries.GetUserByUserName(ctx, username)
}

// ListUsers implements Service.
func (s *service) ListUsers(ctx context.Context, arg db.ListUsersParams) ([]db.User, error) {
	return s.queries.ListUsers(ctx, arg)
}

// UpdateUsers implements Service.
func (s *service) UpdateUsers(ctx context.Context, arg db.UpdateUsersParams) (db.User, error) {
	return s.queries.UpdateUsers(ctx, arg)
}

// UpdateUsersPassword implements Service.
func (s *service) UpdateUsersPassword(ctx context.Context, arg db.UpdateUsersPasswordParams) error {
	return s.queries.UpdateUsersPassword(ctx, arg)
}
