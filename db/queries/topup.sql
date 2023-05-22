-- name: CreateTopUps :one
INSERT INTO topups (
    user_id,
    wallet_id,
    amount,
    topup_date,
    description

) VALUES (
    $1, $2, $3, $4, $5
) RETURNINg *;