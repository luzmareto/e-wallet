-- name: CreateTransaction :exec
INSERT INTO transactions (
    user_id,
    wallet_id,
    amount, 
    description,
    transaction_type
) VALUES (
    $1, $2, $3, $4, $5
);

-- name: GetTransactionUserID :many
SELECT * FROM transactions WHERE user_id = $1;

-- name: GetTransactionWalletID :many
SELECT * FROM transactions WHERE wallet_id = $1;