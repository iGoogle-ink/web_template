package model

import "time"

type Teacher struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Subject string    `json:"subject"`
	Ctime   time.Time `json:"ctime"`
	Mtime   time.Time `json:"mtime"`
}

type TeacherListRsp struct {
	TeacherList []*Teacher `json:"teacher_list"`
}
