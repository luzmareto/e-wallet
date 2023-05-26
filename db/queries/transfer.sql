-- name: CreateTransfers :one
INSERT INTO Transfers (
    from_wallet_id,
    to_wallet_id,
    amount,
    description
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetTransfersByFromWalletID :many
SELECT * FROM transfers 
WHERE from_wallet_id = $1;

-- name: GetTransfersByFromWalletIdAndToWalletId :many
SELECT * FROM transfers 
WHERE from_wallet_id = $1
AND to_wallet_id = $2;