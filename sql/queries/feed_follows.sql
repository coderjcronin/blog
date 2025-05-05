-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2)
    RETURNING *
)
SELECT 
    inserted_feed_follow.*, feeds.name AS feed_name, users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON inserted_feed_follow.user_id=users.id
INNER JOIN feeds ON inserted_feed_follow.feed_id=feeds.id;

-- name: GetFeedsFollowing :many
SELECT feed_follows.*, feeds.name AS feed_name FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id=feeds.id
WHERE feed_follows.user_id=$1;

-- name: DeleteFollowByUrl :exec
DELETE FROM feed_follows o
USING feeds f
WHERE f.id = o.feed_ID AND f.url=$1 AND o.user_id=$2;