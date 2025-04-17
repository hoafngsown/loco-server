-- +goose Up
Create table users (
  id uuid primary key,
  email varchar(255) not null,
  password varchar(255) not null,
  display_name varchar(255) not null,
  preferences jsonb not null default '{}',
  
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);

-- +goose Down
Drop table users;
