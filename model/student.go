package model

import "time"

type Student struct {
	Id    int       `json:"id"`
	Name  string    `json:"name"`
	Ctime time.Time `json:"ctime"`
	Mtime time.Time `json:"mtime"`
}

type StudentListRsp struct {
	StudentList []*StudentRsp `json:"student_list"`
}

type StudentRsp struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Ctime int64  `json:"ctime"`
	Mtime int64  `json:"mtime"`
}

func (s *Student) FormatToRsp(sRsp *StudentRsp) {
	sRsp.Id = s.Id
	sRsp.Name = s.Name
	sRsp.Ctime = s.Ctime.Unix()
	sRsp.Mtime = s.Mtime.Unix()
}
