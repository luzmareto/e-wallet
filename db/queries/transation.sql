-- name: CreateTransaction :exec
INSERT INTO transactions (
    user_id,
    wallet_id,
    amount, 
    description
) VALUES (
    $1, $2, $3, $4
);

-- name: GetTransactionUserID :many
SELECT * FROM transactions WHERE user_id = $1;

-- name: GetTransactionWalletID :many
SELECT * FROM transactions WHERE wallet_id = $1;