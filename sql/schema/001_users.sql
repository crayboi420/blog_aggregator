-- +goose Up
CREATE TABLE users(
    id UUID primary key,
    created_at Timestamp not null,
    updated_at Timestamp not null,
    name text not null
);

-- +goose Down
DROP TABLE users;