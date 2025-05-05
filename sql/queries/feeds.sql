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