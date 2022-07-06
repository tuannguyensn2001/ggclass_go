CREATE TABLE posts (
    id serial primary key,
    content varchar,
    created_by int,
    class_id int,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)