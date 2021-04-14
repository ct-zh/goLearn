create schema if not exists testdb;
create table if not exists testdb.user
(
    id         bigint auto_increment primary key,
    username   varchar(100)                        not null,
    password   varchar(100)                        not null,
    email      varchar(100)                        not null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
key(id)
)engine=innodb charset=utf8;
