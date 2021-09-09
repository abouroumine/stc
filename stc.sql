/*CREATE DATABASE stc-users;*/
create table users (
    id serial primary key not null,
    username varchar(255) not null unique,
    password varchar(255) not null,
    role varchar(255) not null
);
insert into users (username, password, role) values ('ayoub', 'ayoub1111', 'Command');
