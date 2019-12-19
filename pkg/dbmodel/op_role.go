package dbmodel

import (
	"time"
)

type OpRole struct {
	Id        int       `json:"id" xorm:"not null pk autoincr comment('自增长ID') INT(11)"`
	Name      string    `json:"name" xorm:"not null default '' comment('名字') VARCHAR(16)"`
	Nickname  string    `json:"nickname" xorm:"not null default '' comment('昵称') VARCHAR(16)"`
	Gender    int       `json:"gender" xorm:"not null default 0 comment('0：未知，1：男，2：女') TINYINT(4)"`
	Reward    string    `json:"reward" xorm:"not null default '' comment('悬赏金') VARCHAR(16)"`
	AbilityId int       `json:"ability_id" xorm:"not null default 0 comment('能力果实id') INT(11)"`
	ShipId    int       `json:"ship_id" xorm:"not null default 0 comment('海贼团id') INT(11)"`
	Address   string    `json:"address" xorm:"not null default '' comment('初始地址') VARCHAR(64)"`
	IsDelete  int       `json:"is_delete" xorm:"not null default 0 comment('0：正常，1：删除') TINYINT(4)"`
	Ctime     time.Time `json:"ctime" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `json:"mtime" xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') TIMESTAMP"`
}
