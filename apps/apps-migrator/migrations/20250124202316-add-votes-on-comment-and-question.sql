-- +migrate Up
ALTER TABLE questions
  ADD COLUMN votes int;

ALTER TABLE comments
  ADD COLUMN votes int;

-- +migrate Down
ALTER TABLE questions
  DROP COLUMN votes;

ALTER TABLE comments
  DROP COLUMN votes;
