CREATE TABLE IF NOT EXISTS users.group (
  group_name  text PRIMARY KEY CHECK (LENGTH(group_name) <= 50),
  UNIQUE (group_name)
);
