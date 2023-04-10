-- +goose Up
CREATE TABLE posts (
  id UUID PRIMARY KEY,
  feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  title TEXT NOT NULL,
  description TEXT,
  url TEXT NOT NULL UNIQUE,
  published_at TIMESTAMP 
);

-- +goose Down
DROP TABLE posts;