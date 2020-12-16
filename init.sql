create schema if not exists dbname;
create table if not exists user.user
(
    user_id         bigint auto_increment primary key,
    user_name   varchar(100)                        not null,
    password   varchar(100)                        not null,
    email      varchar(100)                        not null,
    phone      varchar(100)                        not null,
    created_at timestamp default CURRENT_TIMESTAMP not null
);
