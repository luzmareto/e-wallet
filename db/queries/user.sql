-- name: CreateUsers :one
INSERT INTO users (
    username,
    password,
    email,
    phone_number,
    registration_date
) VALUES (
    $1,$2,$3,$4,$5
) RETURNINg *;