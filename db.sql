create table op_ship
(
    id        int unsigned auto_increment primary key not null comment '自增长id',
    name      varchar(32) default ''                  not null comment '海贼团名称',
    captain   varchar(16) default ''                  not null comment '船长',
    crew_num  int         default 0                   not null comment '船员人数',
    is_delete tinyint     default 0                   not null comment '0：正常，1：删除',
    ctime     timestamp   default CURRENT_TIMESTAMP   not null comment '创建时间',
    mtime     timestamp   default CURRENT_TIMESTAMP   not null on update CURRENT_TIMESTAMP comment '修改时间'
)
    ENGINE = InnoDB
    AUTO_INCREMENT = 1
    DEFAULT CHARSET = utf8mb4
    COMMENT ='海贼团';

create table op_role
(
    id         int unsigned auto_increment primary key not null comment '自增长id',
    name       varchar(16) default ''                  not null comment '名字',
    nickname   varchar(16) default ''                  not null comment '昵称',
    gender     tinyint     default 0                   not null comment '0：未知，1：男，2：女',
    reward     varchar(16) default ''                  not null comment '悬赏金',
    ability_id int         default 0                   not null comment '能力果实id',
    ship_id    int         default 0                   not null comment '海贼团id',
    address    varchar(64) default ''                  not null comment '初始地址',
    is_delete  tinyint     default 0                   not null comment '0：正常，1：删除',
    ctime      timestamp   default CURRENT_TIMESTAMP   not null comment '创建时间',
    mtime      timestamp   default CURRENT_TIMESTAMP   not null on update CURRENT_TIMESTAMP comment '修改时间'
)
    ENGINE = InnoDB
    AUTO_INCREMENT = 1
    DEFAULT CHARSET = utf8mb4
    COMMENT ='海贼王角色表';

create table op_ability
(
    id        int unsigned auto_increment primary key not null comment '自增长id',
    name      varchar(32)  default ''                 not null comment '果实名称',
    introduce varchar(128) default ''                 not null comment '能力介绍',
    is_delete tinyint      default 0                  not null comment '0：正常，1：删除',
    ctime     timestamp    default CURRENT_TIMESTAMP  not null comment '创建时间',
    mtime     timestamp    default CURRENT_TIMESTAMP  not null on update CURRENT_TIMESTAMP comment '修改时间'
)
    ENGINE = InnoDB
    AUTO_INCREMENT = 1
    DEFAULT CHARSET = utf8mb4
    COMMENT ='果实能力';