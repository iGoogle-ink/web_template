package model

import "time"

type Binding struct {
	Id       int       `json:"id" xorm:"'id' autoincr"`
	BindType int       `json:"bind_type" xorm:"'bind_type'"`
	Pid      int       `json:"pid" xorm:"'pid'"`
	Cid      int       `json:"cid" xorm:"'cid'"`
	Ctime    time.Time `json:"ctime" xorm:"'ctime'"`
	Mtime    time.Time `json:"mtime" xorm:"'mtime'"`
}

func (b *Binding) IsZero() bool {
	return b.Pid == 0 && b.Cid == 0
}
