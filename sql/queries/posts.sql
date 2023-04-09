-- name: CreatePost :one
INSERT INTO posts (id, feed_id, created_at, updated_at, title, description, url, published_at, updated_at)
VALUES (
    $1, 
    $2, 
    $3, 
    $4, 
    $5, 
    $6, 
    $7, 
    $8, 
    $9
)
RETURNING *;

