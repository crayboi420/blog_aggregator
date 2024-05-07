-- name: CreateFollow :one
insert into
    feed_follows(id, created_at, updated_at, user_id, feed_id)
values
    ($1, $2, $3, $4, $5) returning *;

-- name: DeleteFollow :exec
delete
from
    feed_follows
where
    id = $1;

-- name: GetFollows :many
select
    *
from
    feed_follows
where
    user_id = $1;