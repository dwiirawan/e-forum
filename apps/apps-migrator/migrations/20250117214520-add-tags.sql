-- +migrate Up
CREATE TABLE tags
(
  id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS tags;
