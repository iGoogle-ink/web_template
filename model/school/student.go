package school

import "time"

type Student struct {
	Id    int       `json:"id" xorm:"autoincr"`
	Name  string    `json:"name" xorm:"'name'"`
	Ctime time.Time `json:"ctime" xorm:"'ctime'"`
	Mtime time.Time `json:"mtime" xorm:"'mtime'"`
}

type StudentAddReq struct {
	Name      string `json:"name"`
	TeacherId int    `json:"teacher_id"`
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

func (s *StudentAddReq) FormatToStudent() (stu *Student) {
	return &Student{
		Name: s.Name,
	}
}
