-- +goose Up
CREATE TABLE users(
        id UUID PRIMARY KEY,
        created_at TIMESTAMP NOT NULL DEFAULT Now(),
        updated_at TIMESTAMP NOT NULL DEFAULT Now(),
        name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE users;
