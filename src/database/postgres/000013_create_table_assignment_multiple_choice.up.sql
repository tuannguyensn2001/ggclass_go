create table assignment_multiple_choice (
    id serial primary key ,
    assignment_id int,
    exercise_multiple_choice_answer_id int,
    answer varchar,
    created_at varchar,
    updated_at varchar,
    deleted_at varchar
)