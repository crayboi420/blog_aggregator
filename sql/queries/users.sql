-- name: CreateUser :one
INSERT INTO
    users (id, created_at, updated_at, name, apikey)
values
    (
        $1,
        $2,
        $3,
        $4,
        encode(sha256(random() :: text :: bytea), 'hex')
    ) RETURNING *;

-- name: GetUserApi :one
Select
    *
from
    users
where
    apikey = $1;