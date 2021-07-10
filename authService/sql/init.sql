create database if not exists authService;
use authService;

create table auth_info
(
    id            bigint auto_increment primary key,
    client_name   varchar(256) not null default '' comment '用户名称',
    client_secret varchar(256) not null default '' comment '客户端秘钥',
    client_id     varchar(256) not null default '' comment '客户端名称',
    callback_url  varchar(512) not null default '' comment '客户端回调地址',
    is_delete     tinyint(1)   not null default 1 comment '是否删除',
    update_time   datetime     not null default current_timestamp,
    create_time   datetime     not null default current_timestamp
);

create table user_info
(
    id          bigint auto_increment primary key,
    username    varchar(128) not null default '',
    password    varchar(128) not null default '',
    salt        varchar(128) not null default '',
    is_delete   tinyint(1)   not null default 1 comment '是否删除',
    update_time datetime     not null default current_timestamp,
    create_time datetime     not null default current_timestamp
);