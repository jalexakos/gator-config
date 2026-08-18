-- psql "postgres://jalexakos:@localhost:5432/gator"
-- +goose Up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    feed_id UUID REFERENCES feeds(id) ON DELETE CASCADE NOT NULL,
    CONSTRAINT user_and_feed UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;
