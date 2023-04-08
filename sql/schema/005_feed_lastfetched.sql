-- +goose UP
ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose DOWN
ALTER TABLE feeds DROP COLUMN last_fetched_at;
