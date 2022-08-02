create table notifications (
    id serial primary key ,
    owner_name varchar,
    owner_avatar varchar,
    created_by int,
    type int,
    type_id int,
    html_content varchar,
    created_at timestamp,
    updated_at timestamp
)