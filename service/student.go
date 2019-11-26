package service

import (
	"log"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"web_template/model"
)

func (s *Service) loadStudent() {
	// mock data
	stu := make(map[int]string)
	for i := 1; i <= 100; i++ {
		stu[i] = "jerry-" + strconv.Itoa(i)
	}
	s.student = stu
}

func (s *Service) loadStudentProc() {
	for {
		time.Sleep(time.Second * 10)
		log.Println("Reload Data Loop")
		s.loadStudent()
	}
}

func (s *Service) StudentList() (rsp *model.StudentRsp, err error) {
	rsp = &model.StudentRsp{
		StudentList: make([]*model.Student, 0),
	}
	for k, v := range s.student {
		rsp.StudentList = append(rsp.StudentList, &model.Student{
			Id:   k,
			Name: v,
		})
	}
	return rsp, nil
}
func (s *Service) StudentById(id int) (rsp *model.Student, err error) {
	if _, ok := s.student[id]; !ok {
		return nil, errors.New("No this Student")
	}
	rsp = &model.Student{
		Id:   id,
		Name: s.student[id],
	}
	return rsp, nil
}
