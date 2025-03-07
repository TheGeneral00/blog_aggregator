-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
        INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
        Values(
                gen_random_uuid(),
                Now(),
                Now(),
                $1,
                $2
        ) RETURNING *
) SELECT 
        inserted_feed_follow.*,
        feeds.name AS feed_name,
        users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds ON inserted_feed_follow.feed_id= feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT 
        feed_follows.*,
        users.name AS user_name,
        feeds.name AS feed_name
FROM feed_follows
INNER JOIN users ON feed_follows.user_id = users.id 
INNER JOIN feeds ON feed_follows.feed_id = feeds.id 
WHERE feed_follows.user_id = $1;

-- name: DeleteFollow :one
DELETE FROM feed_follows 
USING feeds
WHERE feed_follows.user_id = $1
AND feed_follows.feed_id = feeds.id
AND feeds.url = $2
RETURNING (
        SELECT feeds.name
        FROM feeds
        WHERE feeds.url = $2
        LIMIT 1
);
