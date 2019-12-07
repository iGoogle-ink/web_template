package service

import (
	"log"
	"time"

	"web_template/ecode"
	"web_template/model"

	"github.com/pkg/errors"
	"xorm.io/xorm"
)

func (s *Service) loadStudent() {
	stus, err := s.dao.StudentList()
	if err != nil {
		log.Println("Load Students Error :", err)
		return
	}
	s.studentList = stus
	for _, v := range stus {
		s.studentMap[v.Id] = v
	}
}

func (s *Service) loadStudentProc() {
	for {
		time.Sleep(time.Duration(s.c.ReloadTime) * time.Second)
		log.Println("Reload Student Data")
		s.loadStudent()
	}
}

func (s *Service) StudentAdd(req *model.StudentAddReq) (err error) {
	has, err := s.dao.TeacherExistById(req.TeacherId)
	if err != nil {
		return err
	}
	if !has {
		return errors.New("未找到该老师")
	}

	err = s.dao.Transact(func(tx *xorm.Session) error {
		stu := req.FormatToStudent()
		studentId, err := s.dao.TxStudentInsert(tx, stu)
		if err != nil {
			log.Printf("s.dao.TxStudentInsert(%v),Err(%v)\n", stu, err)
			return err
		}
		if err = s.dao.TxAddBinding(tx, model.BindTypeTeacherStudent, req.TeacherId, studentId); err != nil {
			log.Printf("s.dao.TxAddBinding(type:%d,teacherId:%d,studentId:%d),Err(%v)\n", model.BindTypeTeacherStudent, req.TeacherId, studentId, err)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	// todo something other
	return nil
}

func (s *Service) StudentList() (rsp *model.StudentListRsp, err error) {
	if s.studentList == nil {
		s.loadStudent()
	}
	if len(s.studentList) == 0 {
		return nil, ecode.NothingFound
	}
	rsp = new(model.StudentListRsp)
	for _, v := range s.studentList {
		sRsp := new(model.StudentRsp)
		v.FormatToRsp(sRsp)
		rsp.StudentList = append(rsp.StudentList, sRsp)
	}
	return rsp, nil
}

func (s *Service) StudentById(id int) (rsp *model.StudentRsp, err error) {
	if len(s.studentMap) == 0 {
		s.loadStudent()
	}
	if v, ok := s.studentMap[id]; ok {
		rsp = new(model.StudentRsp)
		v.FormatToRsp(rsp)
		return rsp, nil
	}
	return nil, ecode.NothingFound
}
