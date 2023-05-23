-- name: Createwithdrawals :one
INSERT INTO withdrawals (
    user_id,
    wallet_id,
    amount,
    description
) VALUES (
    $1, $2, $3, $4
) RETURNINg *;