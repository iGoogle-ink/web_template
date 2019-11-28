package model

import "time"

type Binding struct {
	Id       int       `json:"id" xorm:"'id'"`
	BindType int       `json:"bind_type" xorm:"'bind_type'"`
	Pid      int       `json:"pid" xorm:"'pid'"`
	Cid      int       `json:"cid" xorm:"'cid'"`
	Ctime    time.Time `json:"ctime" xorm:"'ctime'"`
	Mtime    time.Time `json:"mtime" xorm:"'mtime'"`
}
