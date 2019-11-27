package model

import "time"

type Student struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	CreateTime int64     `json:"ctime" xorm:"-"`
	UpdateTime int64     `json:"mtime" xorm:"-"`
	Ctime      time.Time `json:"-"`
	Mtime      time.Time `json:"-"`
}

type StudentListRsp struct {
	StudentList []*Student `json:"student_list"`
}

func (s *Student) AfterLoad() {
	s.CreateTime = s.Ctime.Unix()
	s.UpdateTime = s.Mtime.Unix()
}
