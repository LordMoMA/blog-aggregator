-- +goose Up
CREATE TABLE posts (
  id UUID PRIMARY KEY,
  feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  url VARCHAR(255) NOT NULL,
  published_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE posts;