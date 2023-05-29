package service

import (
	"context"
	"database/sql"
	"fmt"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

// CreateUsers implements Service.
func (s *service) CreateUsers(ctx context.Context, arg db.CreateUsersParams) (db.User, error) {
	hashedPassword, _ := utils.HashPassword(arg.Password)
	arg.Password = hashedPassword
	user, err := s.store.CreateUsers(ctx, arg)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

// DeleteUsers implements Service.
func (s *service) DeleteUsers(ctx context.Context, id int64) error {
	if _, err := s.store.GetUserById(ctx, id); err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("user with id %d not found", id),
			Err: err,
		}
		return cstErr
	}
	return s.store.DeleteUsers(ctx, id)
}

// GetUserById implements Service.
func (s *service) GetUserById(ctx context.Context, id int64) (db.User, error) {
	var user db.User
	user, err := s.store.GetUserById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			cstErr := &utils.CustomError{
				Msg: fmt.Sprintf("user with id %d not found", id),
				Err: err,
			}
			return user, cstErr
		}
		return user, err
	}
	return user, nil
}

// GetUserByUserName implements Service.
func (s *service) GetUserByUserName(ctx context.Context, username string) (db.User, error) {
	var user db.User
	user, err := s.store.GetUserByUserName(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			cstErr := &utils.CustomError{
				Msg: fmt.Sprintf("user with username %s not found", username),
				Err: err,
			}
			return user, cstErr
		}
		return user, err
	}
	return user, nil
}

// ListUsers implements Service.
func (s *service) ListUsers(ctx context.Context, arg db.ListUsersParams) ([]db.User, error) {
	users, err := s.store.ListUsers(ctx, arg)
	if err != nil {
		return []db.User{}, err
	}
	return users, nil
}

// UpdateUsers implements Service.
func (s *service) UpdateUsers(ctx context.Context, arg db.UpdateUsersParams) (db.User, error) {
	var user db.User
	user, err := s.store.GetUserById(ctx, arg.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			cstErr := &utils.CustomError{
				Msg: fmt.Sprintf("user with id %d not found", arg.ID),
				Err: err,
			}
			return user, cstErr
		}
		return user, err
	}
	user, err = s.store.UpdateUsers(ctx, arg)
	if err != nil {
		return user, err
	}
	return user, nil
}

// UpdateUsersPassword implements Service.
func (s *service) UpdateUsersPassword(ctx context.Context, arg db.UpdateUsersPasswordParams) error {
	if _, err := s.store.GetUserById(ctx, arg.ID); err != nil {
		if err == sql.ErrNoRows {
			cstErr := &utils.CustomError{
				Msg: fmt.Sprintf("user with id %d not found", arg.ID),
				Err: err,
			}
			return cstErr
		}
		return err
	}
	if err := s.store.UpdateUsersPassword(ctx, arg); err != nil {
		return err
	}
	return nil
}

// UpdateUserIDcard implements Service.
func (s *service) UpdateUserIDcard(ctx context.Context, arg db.UpdateUserIDcardParams) error {
	if _, err := s.store.GetUserById(ctx, arg.ID); err != nil {
		if err == sql.ErrNoRows {
			cstErr := &utils.CustomError{
				Msg: fmt.Sprintf("user with id %d not found", arg.ID),
				Err: err,
			}
			return cstErr
		}
		return err
	}
	if err := s.store.UpdateUserIDcard(ctx, arg); err != nil {
		return err
	}
	return nil
}
