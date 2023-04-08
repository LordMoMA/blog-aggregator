-- +goose Up
CREATE TABLE feed_follows (
  id UUID NOT NULL PRIMARY KEY,
  feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;