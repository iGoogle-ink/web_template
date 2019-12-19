package dbmodel

import (
	"time"
)

type Binding struct {
	Id       int       `json:"id" xorm:"not null pk autoincr comment('自增长ID') INT(11)"`
	BindType int       `json:"bind_type" xorm:"not null default 0 comment('绑定关系种类，后期自定义：<1:师生关系(pid:老师,cid:学生)>') TINYINT(4)"`
	Pid      int       `json:"pid" xorm:"not null default 0 comment('被绑定者ID') INT(11)"`
	Cid      int       `json:"cid" xorm:"not null default 0 comment('绑定者ID') INT(11)"`
	IsDelete int       `json:"is_delete" xorm:"not null default 0 comment('0；正常，1：删除') TINYINT(4)"`
	Ctime    time.Time `json:"ctime" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime    time.Time `json:"mtime" xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') TIMESTAMP"`
}
