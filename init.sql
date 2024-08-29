create database if not exists gin_biz_web_api;

use gin_biz_web_api;

create table if not exists `users`
(
    id           bigint(20) unsigned not null auto_increment,
    account      varchar(255) not null default '' comment '账号',
    email        varchar(80) comment '邮箱',
    phone        varchar(40) comment '手机号',
    password     varchar(255) not null default '' comment '密码',
    nickname     varchar(255) not null default '' comment '昵称',
    introduction text comment '自我简介',
    avatar       varchar(255) not null default '' comment '头像地址',
    created_at   int(11) unsigned    not null default 0 comment '创建时间',
    updated_at   int(11) unsigned    not null default 0 comment '更新时间',
    primary key (id),
    unique key unique_email (email),
    unique key unique_phone (phone),
    unique key unique_account (account)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_unicode_ci comment '用户表';