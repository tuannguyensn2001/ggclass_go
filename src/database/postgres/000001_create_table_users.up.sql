CREATE TABLE  users (
    id serial PRIMARY KEY ,
    username VARCHAR   NOT NULL,
    email VARCHAR  UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);