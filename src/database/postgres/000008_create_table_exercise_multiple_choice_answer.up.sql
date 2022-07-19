create table exercise_multiple_choice_answer (
    id serial primary key ,
    exercise_multiple_choice_id int,
    "order" int,
    "type" smallint,
    answer varchar,
    mark float,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)