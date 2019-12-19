CREATE DATABASE `bilibili_comic` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */

create table comic_info
(
    id          int auto_increment comment '自增长id'
        primary key,
    prefix      varchar(8)  default ''                not null comment '表前缀',
    ch_name     varchar(32) default ''                not null comment '中文名',
    origin_name varchar(32) default ''                not null comment '源名称',
    area        varchar(16) default ''                not null comment '国家地区',
    author      varchar(16) default ''                not null comment '作者',
    ep_num      int         default 0                 not null comment '剧集数',
    is_end      tinyint     default 0                 not null comment '1：是，0：否',
    is_delete   tinyint     default 0                 not null comment '0：正常，1：删除',
    ctime       timestamp   default CURRENT_TIMESTAMP not null comment '创建时间',
    mtime       timestamp   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '修改时间'
)
    comment '动漫信息介绍';

create table hs_student
(
    id        int auto_increment comment '自增长ID'
        primary key,
    name      varchar(16) default ''                not null comment '姓名',
    age       tinyint     default 0                 not null comment '年龄',
    gender    tinyint     default 0                 not null comment '0：未知，1：男，2：女',
    is_delete tinyint     default 0                 not null comment '0：正常，1：删除',
    ctime     timestamp   default CURRENT_TIMESTAMP not null comment '创建时间',
    mtime     timestamp   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '修改时间'
)
    comment '高中学生表';

create table hs_teacher
(
    id        int auto_increment comment '自增长ID'
        primary key,
    name      varchar(16) default ''                not null comment '姓名',
    age       tinyint     default 0                 not null comment '年龄',
    gender    tinyint     default 0                 not null comment '0：未知，1：男，2：女',
    subject   varchar(16) default ''                not null comment '学科名称',
    is_delete tinyint     default 0                 not null comment '0：正常，1：删除',
    ctime     timestamp   default CURRENT_TIMESTAMP not null comment '创建时间',
    mtime     timestamp   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '修改时间'
)
    comment '高中老师表';

create table op_ability
(
    id        int auto_increment comment '自增长id'
        primary key,
    name      varchar(32)  default ''                not null comment '果实名称',
    introduce varchar(128) default ''                not null comment '能力介绍',
    is_delete tinyint      default 0                 not null comment '0：正常，1：删除',
    ctime     timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    mtime     timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '修改时间'
)
    comment '果实能力';

create table op_role
(
    id         int auto_increment comment '自增长ID'
        primary key,
    name       varchar(16) default ''                not null comment '名字',
    nickname   varchar(16) default ''                not null comment '昵称',
    gender     tinyint     default 0                 not null comment '0：未知，1：男，2：女',
    reward     varchar(16) default ''                not null comment '悬赏金',
    ability_id int         default 0                 not null comment '能力果实id',
    ship_id    int         default 0                 not null comment '海贼团id',
    address    varchar(64) default ''                not null comment '初始地址',
    is_delete  tinyint     default 0                 not null comment '0：正常，1：删除',
    ctime      timestamp   default CURRENT_TIMESTAMP not null comment '创建时间',
    mtime      timestamp   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '修改时间'
)
    comment '海贼王角色表';

create table op_ship
(
    id        int auto_increment comment '自增长id'
        primary key,
    name      varchar(32) default ''                not null comment '海贼团名称',
    captain   varchar(16) default ''                not null comment '船长',
    crew_num  int         default 0                 not null comment '船员人数',
    is_delete tinyint     default 0                 not null comment '0：正常，1：删除',
    ctime     timestamp   default CURRENT_TIMESTAMP not null comment '创建时间',
    mtime     timestamp   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '修改时间'
)
    comment '海贼团';

