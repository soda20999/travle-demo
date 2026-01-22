create table user
(
    id          bigint auto_increment
        primary key,
    user_id     bigint                              not null,
    username    varchar(64)                         not null,
    password    varchar(64)                         not null,
    email       varchar(64)                         null,
    nickname    varchar(64)                         null,
    avatar_url  varchar(256)                        null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint idx_user_id
        unique (user_id),
    constraint idx_username
        unique (username)
)