create table lessons (
    id serial primary key ,
    name varchar,
    description varchar,
    folder_id int,
    youtube_link varchar,
    created_by int,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)