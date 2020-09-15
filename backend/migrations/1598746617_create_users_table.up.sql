CREATE TABLE IF NOT EXISTS users (
    id bigserial primary key,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    deleted boolean not null default false,
    name text not null,
    email text unique not null
);
