CREATE TABLE exercise_multiple_choice (
    id serial primary key ,
    exercise_id int,
    number_of_questions int,
    mark int,
    file_question_url varchar,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)