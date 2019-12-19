package dbmodel

import (
	"time"
)

type OpAbility struct {
	Id        int       `json:"id" xorm:"not null pk autoincr comment('自增长id') INT(11)"`
	Name      string    `json:"name" xorm:"not null default '' comment('果实名称') VARCHAR(32)"`
	Introduce string    `json:"introduce" xorm:"not null default '' comment('能力介绍') VARCHAR(128)"`
	IsDelete  int       `json:"is_delete" xorm:"not null default 0 comment('0：正常，1：删除') TINYINT(4)"`
	Ctime     time.Time `json:"ctime" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `json:"mtime" xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') TIMESTAMP"`
}
