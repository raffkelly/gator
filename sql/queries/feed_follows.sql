-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name as feed_name,
    users.name as user_name
FROM inserted_feed_follow
INNER JOIN users
ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds
ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, users.name AS user_name, feeds.name AS feed_name
FROM feed_follows
INNER JOIN users
ON feed_follows.user_id = users.id
INNER JOIN feeds
on feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
USING users, feeds
WHERE feed_follows.user_id = users.id
    AND feed_follows.feed_id = feeds.id
    AND users.name = $1
    AND feeds.url = $2;