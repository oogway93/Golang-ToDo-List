CREATE TABLE todo_list
(
    id          SERIAL       NOT NULL UNIQUE,
    title       VARCHAR(256) NOT NULL,
    description VARCHAR(256) NOT NULL
);