package hs

import (
	"time"

	"web_template/pkg/dbmodel"
)

type Teacher struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Gender  int       `json:"gender"`
	Subject string    `json:"subject"`
	Ctime   time.Time `json:"ctime"`
	Mtime   time.Time `json:"mtime"`
}

type TeacherAddReq struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Gender  int    `json:"gender"`
	Subject string `json:"subject"`
}

type TeacherListRsp struct {
	TeacherList []*TeacherRsp `json:"list"`
}

type TeacherRsp struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Gender  int    `json:"gender"`
	Subject string `json:"subject"`
	Ctime   int64  `json:"ctime"`
	Mtime   int64  `json:"mtime"`
}

func (s *Teacher) FormatToRsp(tRsp *TeacherRsp) {
	tRsp.Id = s.Id
	tRsp.Name = s.Name
	tRsp.Age = s.Age
	tRsp.Gender = s.Gender
	tRsp.Subject = s.Subject
	tRsp.Ctime = s.Ctime.Unix()
	tRsp.Mtime = s.Mtime.Unix()
}

func (t *TeacherAddReq) FormatToTeacher() (tch *dbmodel.HsTeacher) {
	return &dbmodel.HsTeacher{
		Name:    t.Name,
		Age:     t.Age,
		Gender:  t.Gender,
		Subject: t.Subject,
	}
}
