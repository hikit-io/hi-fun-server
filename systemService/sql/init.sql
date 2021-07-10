create database if not exists system_service;
use system_service;

create table banner
(
    id   bigint auto_increment primary key,
    path varchar(128) not null default '',
    url  varchar(128) not null default ''
);

insert into banner (path, url)
values ('/index', '');
insert into banner (path, url)
values ('/practice', 'https://ali-cdn.educoder.net/images/avatars/PortalImage/101?t=1622106999');
insert into banner (path, url)
values ('/competition', 'https://ali-cdn.educoder.net/images/avatars/PortalImage/101?t=1622106999');
insert into banner (path, url)
values ('/about', 'https://ali-cdn.educoder.net/images/avatars/PortalImage/101?t=1622106999');


create table carousel
(
    id      bigint auto_increment primary key,
    `order` int          not null default 0,
    url     varchar(512) not null default '',
    img     varchar(512) not null default '',
    is_show tinyint      not null default 0
);
