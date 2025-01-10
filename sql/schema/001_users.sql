-- +goose Up
CREATE TABLE users(
        id UUID PRIMARY KEY,
        created_at TIMESTAMP NOT NULL DEFAULT Now(),
        updated_at TIMESTAMP NOT NULL DEFAULT Now(),
        name VARCHAR(255) NOT NULL,
        UNIQUE(name)
);

-- +goose Down
DROP TABLE users;
