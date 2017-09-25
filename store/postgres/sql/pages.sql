DROP TABLE IF EXISTS pages CASCADE;
CREATE TABLE pages (
  id          SERIAL PRIMARY KEY,
  type        TEXT,
  placement   SMALLINT,
  doc_id      TEXT UNIQUE NOT NULL,
  created     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  title       TEXT,
  slug        TEXT NOT NULL,
  md          TEXT
);

CREATE INDEX slug_idx ON pages (slug);
