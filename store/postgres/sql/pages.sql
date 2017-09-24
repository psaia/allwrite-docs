DROP TABLE IF EXISTS pages CASCADE;
CREATE TABLE pages (
  id                SERIAL PRIMARY KEY,
  doc_id            text UNIQUE NOT NULL,
  created           timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated           timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  title             text,
  slug              text UNIQUE NOT NULL,
  md                text
);
