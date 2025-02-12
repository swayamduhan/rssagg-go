-- name: AddFeed :one
INSERT INTO feeds (id, name, url)
VALUES (uuid_generate_v4(), $1, $2)
RETURNING *;

-- name: DeleteFeed :exec
DELETE FROM feeds WHERE id = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1;

-- name: MarkFeedFetched :one
UPDATE feeds SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;