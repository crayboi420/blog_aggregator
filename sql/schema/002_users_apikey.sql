-- +goose Up
Alter table
    users
Add
    apikey varchar(64) not null unique 
    default encode(sha256(random() :: text :: bytea), 'hex');

-- +goose Down
Alter table
    users remove column apikey;