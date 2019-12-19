package hs

import (
	"time"

	"web_template/pkg/dbmodel"
)

type Student struct {
	Id     int       `json:"id"`
	Name   string    `json:"name"`
	Age    int       `json:"age"`
	Gender int       `json:"gender"`
	Ctime  time.Time `json:"ctime"`
	Mtime  time.Time `json:"mtime"`
}

type StudentAddReq struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Gender    int    `json:"gender"`
	TeacherId int    `json:"teacher_id"`
}

type StudentListRsp struct {
	StudentList []*StudentRsp `json:"list"`
}

type StudentRsp struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
	Ctime  int64  `json:"ctime"`
	Mtime  int64  `json:"mtime"`
}

func (s *Student) FormatToRsp(sRsp *StudentRsp) {
	sRsp.Id = s.Id
	sRsp.Name = s.Name
	sRsp.Age = s.Age
	sRsp.Gender = s.Gender
	sRsp.Ctime = s.Ctime.Unix()
	sRsp.Mtime = s.Mtime.Unix()
}

func (s *StudentAddReq) FormatToStudent() (stu *dbmodel.HsStudent) {
	return &dbmodel.HsStudent{
		Name:   s.Name,
		Age:    s.Age,
		Gender: s.Gender,
	}
}
