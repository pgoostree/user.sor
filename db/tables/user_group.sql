CREATE TABLE IF NOT EXISTS users.user_group (
  user_id     text NOT NULL REFERENCES users.user(user_id) ON DELETE CASCADE,
  group_name  text NOT NULL REFERENCES users.group(group_name) ON DELETE CASCADE,
  UNIQUE (user_id, group_name)
);

CREATE INDEX idx_group ON users.user_group (group_name)