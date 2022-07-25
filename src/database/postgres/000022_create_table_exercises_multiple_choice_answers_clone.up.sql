create table exercises_multiple_choice_answers_clone (
    id serial primary key ,
    exercise_multiple_choice_id int,
    "order" int,
    type smallint,
    answer varchar,
    mark int,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)