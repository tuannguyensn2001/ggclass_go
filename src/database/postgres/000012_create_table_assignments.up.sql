create table assignments (
    id serial primary key ,
    exercise_id int,
    time_late int,
    user_id int,
    mark float,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)