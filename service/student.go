package service

import (
	"log"
	"time"

	"web_template/ecode"
	"web_template/model"
)

func (s *Service) loadStudent() {
	stus, err := s.dao.StudentList()
	if err != nil {
		log.Println("Load Students Error :", err)
		return
	}
	s.studentList = stus
	for _, v := range stus {
		s.studentMap[v.Id] = v.Name
	}
}

func (s *Service) loadStudentProc() {
	for {
		time.Sleep(time.Duration(s.c.ReloadTime) * time.Second)
		log.Println("Reload Student Data")
		s.loadStudent()
	}
}

func (s *Service) StudentList() (rsp *model.StudentListRsp, err error) {
	if s.studentList == nil {
		s.loadStudent()
	}
	if len(s.studentList) == 0 {
		return nil, ecode.NothingFound
	}
	rsp = &model.StudentListRsp{
		StudentList: s.studentList,
	}
	return rsp, nil
}
func (s *Service) StudentById(id int) (rsp *model.Student, err error) {
	if len(s.studentMap) == 0 {
		s.loadStudent()
	}
	if v, ok := s.studentMap[id]; ok {
		rsp = &model.Student{
			Id:   id,
			Name: v,
		}
		return rsp, nil
	}
	return nil, ecode.NothingFound
}
