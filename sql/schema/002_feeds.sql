-- +goose Up
CREATE TABLE feeds (
        id UUID PRIMARY KEY,
        name varchar(255) NOT NULL,
        url varchar(255) UNIQUE,
        created_at TIMESTAMP NOT NULL DEFAULT Now(),
        updated_at TIMESTAMP NOT NULL DEFAULT Now(),
        user_id UUID,
        CONSTRAINT fk_user_id
        FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
