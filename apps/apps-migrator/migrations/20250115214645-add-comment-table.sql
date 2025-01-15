-- +migrate Up
CREATE TABLE IF NOT EXISTS comments
(
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  content     TEXT      NOT NULL,
  parent_type SMALLINT  NOT NULL,
  parent_id   UUID      NOT NULL,
  user_id     UUID      NOT NULL,
  created_at  TIMESTAMP NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id)
);


-- +migrate Down
DROP TABLE IF EXISTS comments;
