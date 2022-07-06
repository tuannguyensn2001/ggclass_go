CREATE TABLE classes (
    id serial PRIMARY KEY ,
    name VARCHAR NOT NULL,
    description VARCHAR,
    room VARCHAR,
    topic VARCHAR,
    code VARCHAR,
    created_by INT,
    updated_by INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)