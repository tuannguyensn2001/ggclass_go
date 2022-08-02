create table notification_to_user(
    id serial primary key ,
    notification_id int,
    user_id int,
    seen int,
    created_at timestamp,
    updated_at timestamp
)