// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"
)

type Querier interface {
	CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error)
	DeleteUsers(ctx context.Context, id int32) error
	GetUserById(ctx context.Context, id int32) (User, error)
	GetUserByUserName(ctx context.Context, username string) (User, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateUsers(ctx context.Context, arg UpdateUsersParams) (User, error)
	UpdateUsersPassword(ctx context.Context, arg UpdateUsersPasswordParams) error
}

var _ Querier = (*Queries)(nil)
