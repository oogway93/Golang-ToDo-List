CREATE TABLE todo_list
(
    id          SERIAL       NOT NULL UNIQUE,
    title       VARCHAR(256) NOT NULL,
    description VARCHAR(256) NOT NULL
);

CREATE TABLE todo_item
(
    id          SERIAL       NOT NULL UNIQUE,
    title       VARCHAR(256) NOT NULL,
    description VARCHAR(256) NOT NULL,
    done        Bool         NOT NULL DEFAULT FALSE
);

CREATE TABLE todo_list_items
(
    id      SERIAL                                          NOT NULL UNIQUE,
    list_id INT REFERENCES todo_list (id) ON DELETE CASCADE NOT NULL,
    item_id INT REFERENCES todo_item (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE users
(
    id       SERIAL       NOT NULL UNIQUE,
    username VARCHAR(64)  NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL
);
