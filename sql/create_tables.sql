CREATE TABLE users
(
    id            serial       PRIMARY KEY,
    login      varchar(255) not null unique,
    password   varchar(255) not null
);

CREATE TABLE note
(
    id          serial       not null unique,
    description varchar(255) not null,
    user_Id int references users (id) not null
);