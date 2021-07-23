CREATE TABLE task(
     id SERIAL PRIMARY KEY NOT NULL, task text,
     status INTEGER DEFAULT 0
);