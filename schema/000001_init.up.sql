CREATE TABLE person_info (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    age INTEGER NOT NULL,
    gender VARCHAR(10) NOT NULL,
    nationalize VARCHAR(10) NOT NULL
)