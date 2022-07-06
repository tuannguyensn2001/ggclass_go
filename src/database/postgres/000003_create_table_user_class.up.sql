CREATE TABLE user_class (
    id serial primary key,
    user_id int not null,
    class_id int not null,
    role int default 2,
    status int default 2,
    created_at timestamp,
    updated_at timestamp
)