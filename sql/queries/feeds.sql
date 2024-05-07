-- name: CreateFeed :one
insert into
    feeds(id, created_at, updated_at, name, url, user_id)
values
    ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetFeeds :many
select
    *
from
    feeds;

-- name: GetNextFeedsToFetch :many
select
    *
from
    feeds
order by
    last_fetched_at is null,
    last_fetched_at asc
limit
    $1;

-- name: MarkFeedFetched :one
update
    feeds
set
    updated_at = $1,
    last_fetched_at = $1
where
    id = $2
RETURNING *;