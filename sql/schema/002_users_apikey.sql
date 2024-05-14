-- +goose Up
ALTER TABLE users
    ADD apikey varchar(64) NOT NULL UNIQUE DEFAULT encode(sha256(random()::text::bytea), 'hex');

-- +goose Down
ALTER TABLE users
    DROP COLUMN apikey;

