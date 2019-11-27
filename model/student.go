package model

import "time"

type Student struct {
	Id    int       `json:"id"`
	Name  string    `json:"name"`
	Ctime time.Time `json:"ctime"`
	Mtime time.Time `json:"mtime"`
}

type StudentListRsp struct {
	StudentList []*Student `json:"student_list"`
}

func (s *Student) TimeToUnix() {

}
