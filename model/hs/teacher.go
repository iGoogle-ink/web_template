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

func (t *Teacher) FormatToRsp() (tRsp *TeacherRsp) {
	tRsp = new(TeacherRsp)
	tRsp.Id = t.Id
	tRsp.Name = t.Name
	tRsp.Age = t.Age
	tRsp.Gender = t.Gender
	tRsp.Subject = t.Subject
	tRsp.Ctime = t.Ctime.Unix()
	tRsp.Mtime = t.Mtime.Unix()
	return tRsp
}

func (t *TeacherAddReq) FormatToTeacher() (tch *dbmodel.HsTeacher) {
	return &dbmodel.HsTeacher{
		Name:    t.Name,
		Age:     t.Age,
		Gender:  t.Gender,
		Subject: t.Subject,
	}
}

func (t *Teacher) FormatToMap() map[string]interface{} {
	mp := make(map[string]interface{})
	mp["id"] = t.Id
	mp["name"] = t.Name
	mp["age"] = t.Age
	mp["gender"] = t.Gender
	mp["subject"] = t.Subject
	mp["ctime"] = t.Ctime.Unix()
	mp["mtime"] = t.Mtime.Unix()
	return mp
}
