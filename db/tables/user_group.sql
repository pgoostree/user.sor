CREATE TABLE IF NOT EXISTS users.user_group (
  user_id     text REFERENCES users.user ON DELETE CASCADE,
  group_name  text REFERENCES users.group ON DELETE CASCADE,
  UNIQUE (user_id, group_name)
);

CREATE INDEX idx_group ON users.user_group (group_name)