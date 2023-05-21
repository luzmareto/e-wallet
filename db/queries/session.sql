-- name: CreateSession :one
INSERT INTO sessions (
    id,
    username,
    refresh_token,
    user_agent,
    client_ip,
    is_blocked,
    expired_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNINg *;

-- name: GetSessions :one
SELECT * FROM sessions
WHERE id = $1 LIMIT 1;