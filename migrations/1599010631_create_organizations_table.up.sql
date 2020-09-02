CREATE TABLE IF NOT EXISTS organizations (
    id bigserial primary key,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    name text not null
);
