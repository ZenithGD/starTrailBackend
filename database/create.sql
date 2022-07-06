CREATE TABLE users (
    nickname        varchar         primary key,
    descr           text,
    email           varchar,
    password        varchar(100)    not null
);
