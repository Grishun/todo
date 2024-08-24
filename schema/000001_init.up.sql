CREATE TABLE users
(
    "id"            SERIAL       NOT NULL UNIQUE,
    "name"          VARCHAR(255) NOT NULL,
    "username"      varchar(255) not null unique,
    "password_hash" varchar(255) not null
);

CREATE TABLE todo_lists
(
    id          serial       not null unique,
    title      varchar(255) not null,
    description varchar(255) not null
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_items
(
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL
);

CREATE TABLE lists_items
(
    id SERIAL NOT NULL UNIQUE,
    item_id INT REFERENCES todo_items (id) ON DELETE CASCADE      NOT NULL,
    list_id INT REFERENCES todo_lists (id) ON DELETE CASCADE NOT NULL
);
