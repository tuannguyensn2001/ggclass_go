create table notification_from_teacher_to_class (
    id serial primary key ,
    class_id int,
    content varchar,
    created_by int,
    created_at timestamp,
    updated_at timestamp
)