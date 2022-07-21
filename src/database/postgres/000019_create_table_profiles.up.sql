create table profiles
(
    id         serial primary key,
    user_id    int,
    avatar     varchar,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)