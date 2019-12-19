package dbmodel

import (
	"time"
)

type ComicInfo struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('自增长id') INT(11)"`
	Prefix     string    `json:"prefix" xorm:"not null default '' comment('表前缀') VARCHAR(8)"`
	ChName     string    `json:"ch_name" xorm:"not null default '' comment('中文名') VARCHAR(32)"`
	OriginName string    `json:"origin_name" xorm:"not null default '' comment('源名称') VARCHAR(32)"`
	Area       string    `json:"area" xorm:"not null default '' comment('国家地区') VARCHAR(16)"`
	Author     string    `json:"author" xorm:"not null default '' comment('作者') VARCHAR(16)"`
	EpNum      int       `json:"ep_num" xorm:"not null default 0 comment('剧集数') INT(11)"`
	IsEnd      int       `json:"is_end" xorm:"not null default 0 comment('1：是，0：否') TINYINT(4)"`
	IsDelete   int       `json:"is_delete" xorm:"not null default 0 comment('0：正常，1：删除') TINYINT(4)"`
	Ctime      time.Time `json:"ctime" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime      time.Time `json:"mtime" xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') TIMESTAMP"`
}
