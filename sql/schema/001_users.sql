-- +goose Up
CREATE TABLE users (
  id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE users;