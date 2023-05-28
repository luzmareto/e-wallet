-- name: CreateUsers :one
INSERT INTO users (
    username,
    password,
    email,
    phone_number
) VALUES (
    $1,$2,$3,$4
) RETURNINg *;

-- name: DeleteUsers :exec
DELETE FROM users WHERE id = $1;

-- name: UpdateUsers :one
UPDATE users SET email = $2, phone_number = $3 WHERE id = $1 RETURNINg *;

-- name: UpdateUsersPassword :exec
UPDATE users SET password = $2 WHERE id = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByUserName :one
SELECT * FROM users WHERE username = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateUserIDcard :exec
UPDATE users SET id_card = $2 WHERE id = $1;