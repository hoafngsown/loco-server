-- +goose Up
CREATE TABLE preference_metadata (
  id uuid primary key,
  key varchar(255) not null unique,
  type varchar(50) not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);

-- +goose Down
DROP TABLE preference_metadata;
