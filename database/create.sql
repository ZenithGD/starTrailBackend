CREATE TABLE users (
    id              serial          primary key,
    nickname        varchar         not null,
    descr           text,
    email           varchar,
    password        varchar(100)    not null
);
