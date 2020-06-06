CREATE TABLE IF NOT EXISTS users.user (
  user_id             text PRIMARY KEY CHECK (LENGTH(user_id) <= 50),
  first_name          text NOT NULL CHECK (LENGTH(first_name) <= 50),
  last_name           text NOT NULL CHECK (LENGTH(last_name) <= 50),
  UNIQUE (user_id)
);

CREATE INDEX idx_user_id ON users.user (user_id)