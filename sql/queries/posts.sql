-- name: CreatePost :exec
INSERT INTO posts(id, created_at, updated_at, title, url, description, published_at, feed_id)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (url)
    DO UPDATE SET
        updated_at = $3;

-- name: GetPostsByUser :many
SELECT
    *
FROM
    posts
WHERE
    feed_id IN (
        SELECT
            id
        FROM
            feeds
        WHERE
            user_id = $1)
ORDER BY
    published_at DESC
LIMIT $2;

