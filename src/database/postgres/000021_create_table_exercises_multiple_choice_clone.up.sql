create table exercises_multiple_choice_clone (
    id serial primary key ,
    number_of_question int,
    mark int,
    file_question_url  varchar,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)