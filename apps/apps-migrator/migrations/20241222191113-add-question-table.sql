-- +migrate Up
CREATE TABLE IF NOT EXISTS questions
(
  id         UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
  title      VARCHAR(255) NOT NULL,
  content    TEXT         NOT NULL,
  user_id    UUID         NOT NULL,
  views      INT          NOT NULL DEFAULT 0,
  created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users (id)
);

-- +migrate Down
DROP TABLE IF EXISTS questions;
