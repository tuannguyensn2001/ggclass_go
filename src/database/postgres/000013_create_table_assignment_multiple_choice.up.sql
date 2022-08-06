create table assignment_multiple_choice (
    id serial primary key ,
    assignment_id int,
    exercise_multiple_choice_answer_id int,
    answer varchar,
    created_at timestamp,
    updated_at timestamp,
    deleted_at varchar
)