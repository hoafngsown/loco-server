-- +goose Up
CREATE TYPE preference_type AS ENUM ('vibe', 'style');

CREATE TABLE preference_metadata (
  id uuid primary key,
  key varchar(255) not null unique,
  type preference_type not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);

-- +goose Down
DROP TABLE preference_metadata;
DROP TYPE preference_type;
