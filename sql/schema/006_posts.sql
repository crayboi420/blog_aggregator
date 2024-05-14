-- +goose Up
CREATE TABLE posts(
    id uuid PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    title varchar(256),
    url varchar(256) NOT NULL UNIQUE,
    description text,
    published_at timestamp,
    feed_id uuid NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;

