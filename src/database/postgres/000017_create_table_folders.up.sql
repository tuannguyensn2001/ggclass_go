create table folders (
    id serial primary key ,
    name varchar,
    class_id int,
    created_by int,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)