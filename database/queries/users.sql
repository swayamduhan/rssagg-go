-- name: CreateUser :one
INSERT INTO users (id, name)
VALUES (uuid_generate_v4(), $1)
RETURNING *;


-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;