package dbmodel

import (
	"time"
)

type OpShip struct {
	Id       int       `json:"id" xorm:"not null pk autoincr comment('自增长id') INT(11)"`
	Name     string    `json:"name" xorm:"not null default '' comment('海贼团名称') VARCHAR(32)"`
	Captain  string    `json:"captain" xorm:"not null default '' comment('船长') VARCHAR(16)"`
	CrewNum  int       `json:"crew_num" xorm:"not null default 0 comment('船员人数') INT(11)"`
	IsDelete int       `json:"is_delete" xorm:"not null default 0 comment('0：正常，1：删除') TINYINT(4)"`
	Ctime    time.Time `json:"ctime" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime    time.Time `json:"mtime" xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') TIMESTAMP"`
}
