// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package db

import (
	"context"
)

const createUsers = `-- name: CreateUsers :one
INSERT INTO users (
    username,
    password,
    email,
    phone_number
) VALUES (
    $1,$2,$3,$4
) RETURNINg id, role, username, password, email, phone_number, id_card, registration_date
`

type CreateUsersParams struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUsers,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.PhoneNumber,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.PhoneNumber,
		&i.IDCard,
		&i.RegistrationDate,
	)
	return i, err
}

const deleteUsers = `-- name: DeleteUsers :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUsers(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUsers, id)
	return err
}

const getUserById = `-- name: GetUserById :one
SELECT id, role, username, password, email, phone_number, id_card, registration_date FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.PhoneNumber,
		&i.IDCard,
		&i.RegistrationDate,
	)
	return i, err
}

const getUserByUserName = `-- name: GetUserByUserName :one
SELECT id, role, username, password, email, phone_number, id_card, registration_date FROM users WHERE username = $1
`

func (q *Queries) GetUserByUserName(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUserName, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.PhoneNumber,
		&i.IDCard,
		&i.RegistrationDate,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, role, username, password, email, phone_number, id_card, registration_date FROM users ORDER BY id LIMIT $1 OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Role,
			&i.Username,
			&i.Password,
			&i.Email,
			&i.PhoneNumber,
			&i.IDCard,
			&i.RegistrationDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUsers = `-- name: UpdateUsers :one
UPDATE users SET email = $2, phone_number = $3 WHERE id = $1 RETURNINg id, role, username, password, email, phone_number, id_card, registration_date
`

type UpdateUsersParams struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (q *Queries) UpdateUsers(ctx context.Context, arg UpdateUsersParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUsers, arg.ID, arg.Email, arg.PhoneNumber)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.PhoneNumber,
		&i.IDCard,
		&i.RegistrationDate,
	)
	return i, err
}

const updateUsersPassword = `-- name: UpdateUsersPassword :exec
UPDATE users SET password = $2 WHERE id = $1
`

type UpdateUsersPasswordParams struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
}

func (q *Queries) UpdateUsersPassword(ctx context.Context, arg UpdateUsersPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateUsersPassword, arg.ID, arg.Password)
	return err
}
