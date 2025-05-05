-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: ListFeedsWithCreators :many
SELECT feeds.name, feeds.url, users.name
FROM feeds
INNER JOIN users ON feeds.user_id=users.id;

-- name: LookupFeedByUrl :one
SELECT name, id FROM feeds WHERE url=$1;

-- name: MarkFeedFetched :exec
UPDATE feeds SET updated_at=NOW(), last_fetched_at=NOW() WHERE id=$1;

-- name: GetNextFeedToFetch :one
SELECT feeds.*
FROM feeds
WHERE id IN (
    SELECT feed_follows.feed_id
    FROM feed_follows
    WHERE feed_follows.user_id=$1
)
ORDER BY feeds.last_fetched_at ASC NULLS FIRST LIMIT 1;