-- name: CreatePost :one
INSERT INTO posts (id, feed_id, created_at, updated_at, title, description, url, published_at)
VALUES (
    $1, 
    $2, 
    $3, 
    $4, 
    $5, 
    $6, 
    $7, 
    $8
)
RETURNING *;
--

-- name: GetPostsByUser :many
SELECT posts.* FROM posts
JOIN feed_follows ON feed_follows.feed_id = posts.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;
--

