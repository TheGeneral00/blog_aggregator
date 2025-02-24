-- name: AddFeed :one
INSERT INTO feeds (id, name, url, created_at, updated_at, user_id)
Values(
        gen_random_uuid(),
        $1,
        $2,
        Now(),
        Now(),
        $3
) RETURNING *;
