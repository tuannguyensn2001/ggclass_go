CREATE TABLE comments (
    id serial primary key,
    content varchar,
    post_id int,
    created_by int,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)