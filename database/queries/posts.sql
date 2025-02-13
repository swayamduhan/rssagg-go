-- name: CreatePost :one
INSERT INTO posts (feed_id, title, description, link, published_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.* FROM posts
JOIN followed_feeds ON followed_feeds.feed_id = posts.feed_id
WHERE followed_feeds.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;