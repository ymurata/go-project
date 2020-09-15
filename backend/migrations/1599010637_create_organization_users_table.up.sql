CREATE TABLE IF NOT EXISTS organization_users (
    id bigserial primary key,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    organization_id bigint not null references organizations (id),
    user_id bigint not null unique references users (id)
);
