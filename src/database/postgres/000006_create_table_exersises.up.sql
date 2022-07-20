CREATE TABLE exercise (
    id serial primary key ,
    name varchar,
    password varchar,
    time_to_do int,
    time_start timestamp,
    time_end timestamp,
    is_test smallint,
    preview_view_question smallint,
    role_student smallint,
    number_of_time_to_do smallint,
    mode smallint,

    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)