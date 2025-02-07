-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
  id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name       VARCHAR(255)        NOT NULL,
  username   VARCHAR(255) UNIQUE NOT NULL,
  email      VARCHAR(255) UNIQUE NOT NULL,
  picture    VARCHAR(255),
  hash       VARCHAR(255)        NOT NULL,
  salt       VARCHAR(255)        NOT NULL,
  is_active  BOOLEAN          DEFAULT true,
  created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS users;
