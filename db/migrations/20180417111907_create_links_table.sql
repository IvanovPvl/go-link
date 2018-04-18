-- migrate:up

CREATE TABLE links (
  id SERIAL NOT NULL CONSTRAINT links_pkey PRIMARY KEY,
  url TEXT NOT NULL,
  short VARCHAR(6) NOT NULL CONSTRAINT links_short_unique UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);

CREATE TABLE stats (
  id SERIAL NOT NULL CONSTRAINT stats_pkey PRIMARY KEY,
  referer TEXT DEFAULT NULL,
  user_agent VARCHAR(255) NOT NULL,
  ip CIDR NOT NULL,
  link_id INTEGER NOT NULL
    CONSTRAINT stats_link_id_fk
    REFERENCES links (id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);

CREATE INDEX stats_link_id_idx ON stats (link_id);

-- migrate:down

DROP TABLE stats;
DROP TABLE links;
