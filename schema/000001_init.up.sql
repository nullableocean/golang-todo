CREATE TABLE users
(
    id serial not null unique,
    username varchar(255) not null unique,
    name varchar(255) not null,
    password varchar(255) not null
);

CREATE TABLE todo_lists
(
    id serial not null unique,
    title VARCHAR(255) not null,
    description TEXT
);

CREATE TABLE tasks
(
    id serial not null unique,
    title VARCHAR(255) not null ,
    description TEXT,
    is_done boolean not null default false
);

CREATE TABLE users_lists
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    list_id int references todo_lists(id) on delete cascade not null
);

CREATE TABLE tasks_lists
(
    id serial not null unique,
    task_id int references tasks(id) on delete cascade not null,
    list_id int references todo_lists(id) on delete cascade not null
);