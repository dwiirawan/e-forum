-- +migrate Up
CREATE TABLE votes
(
  id     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  parent_type SMALLINT NOT NULL,
  parent_id   UUID     NOT NULL,
  user_id     UUID     NOT NULL,
  vote_type   SMALLINT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id)

);


-- +migrate Down
DROP TABLE IF EXISTS votes;
