-- name: CreateMerchants :one
INSERT INTO merchants (
    merchant_name,
    description,
    website,
    address
) VALUES (
    $1, $2, $3, $4
) RETURNINg *;

-- name: DeleteMerchants :exec
DELETE FROM merchants WHERE id = $1;

-- name: GetMerchantsById :one
SELECT * FROM merchants WHERE id = $1;

-- name: ListMerchants :many
SELECT * FROM merchants LIMIT $1 OFFSET $2 ;

-- name: GetMerchantsByMerchantsName :one
SELECT * FROM merchants WHERE merchant_name = $1;

-- name: UpdatMerchants :one
UPDATE merchants SET description = $2, address = $3 WHERE id = $1 RETURNINg *;

-- name: AddMerchantBalance :one
UPDATE merchants SET balance = balance + $2 WHERE id = $1 RETURNINg *;

