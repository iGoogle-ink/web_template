package model

import "time"

type Teacher struct {
	Id      int       `json:"id" xorm:"autoincr"`
	Name    string    `json:"name" xorm:"'name'"`
	Subject string    `json:"subject" xorm:"'subject'"`
	Ctime   time.Time `json:"ctime" xorm:"'ctime'"`
	Mtime   time.Time `json:"mtime" xorm:"'mtime'"`
}

type TeacherAddReq struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
}

type TeacherListRsp struct {
	TeacherList []*TeacherRsp `json:"teacher_list"`
}

type TeacherRsp struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Ctime   int64  `json:"ctime"`
	Mtime   int64  `json:"mtime"`
}

func (s *Teacher) FormatToRsp(tRsp *TeacherRsp) {
	tRsp.Id = s.Id
	tRsp.Name = s.Name
	tRsp.Subject = s.Subject
	tRsp.Ctime = s.Ctime.Unix()
	tRsp.Mtime = s.Mtime.Unix()
}

func (t *TeacherAddReq) FormatToTeacher() (tch *Teacher) {
	return &Teacher{
		Name:    t.Name,
		Subject: t.Subject,
	}
}
