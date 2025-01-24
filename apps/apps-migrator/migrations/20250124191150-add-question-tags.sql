-- +migrate Up
CREATE TABLE IF NOT EXISTS question_tags
(
  question_id UUID NOT NULL,
  tag_id           UUID NOT NULL,
  PRIMARY KEY (question_id, tag_id),
  FOREIGN KEY (question_id) REFERENCES questions (id) ON DELETE CASCADE,
  FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE
);


-- +migrate Down
DROP TABLE IF EXISTS question_tags;
