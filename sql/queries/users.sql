-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, name, apikey)
    VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
ON CONFLICT (name)
    DO UPDATE SET
        updated_at = $3
    RETURNING
        *;

-- name: GetUserApi :one
SELECT
    *
FROM
    users
WHERE
    apikey = $1;

