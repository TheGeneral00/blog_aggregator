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

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByURL :one
Select * FROM feeds
Where url = $1;

-- name: GetFeedByID :one
SELECT * FROM feeds
WHERE id = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds 
SET last_fetched_at = CURRENT_TIMESTAMP AT TIME ZONE 'UTC', updated_at = CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
WHERE feeds.id = $1;

-- name: GetNextFeedToFetch :one
SELECT feeds.* FROM feeds
JOIN feed_follows ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY last_fetched_at ASC NULLS FIRST 
LIMIT 1;
