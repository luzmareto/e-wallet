-- name: CreateWallets :one
INSERT INTO wallets (
    user_id,
    balance,
    currency
) VALUES (
    $1, $2, $3
) RETURNINg *;

-- name: AddWalletBalance :one
UPDATE wallets 
SET balance = balance + $2
WHERE id = $1 RETURNINg *;

-- name: GetWalletById :one
SELECT * FROM wallets WHERE id = $1;