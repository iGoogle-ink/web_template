package dbmodel

import (
	"time"
)

type HsStudent struct {
	Id       int       `json:"id" xorm:"not null pk autoincr comment('自增长ID') INT(11)"`
	Name     string    `json:"name" xorm:"not null default '' comment('姓名') VARCHAR(16)"`
	Age      int       `json:"age" xorm:"not null default 0 comment('年龄') TINYINT(4)"`
	Gender   int       `json:"gender" xorm:"not null default 0 comment('0：未知，1：男，2：女') TINYINT(4)"`
	IsDelete int       `json:"is_delete" xorm:"not null default 0 comment('0：正常，1：删除') TINYINT(4)"`
	Ctime    time.Time `json:"ctime" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime    time.Time `json:"mtime" xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') TIMESTAMP"`
}
