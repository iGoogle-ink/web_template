package hs

import (
	"log"
	"time"

	"web_template/ecode"
	"web_template/model"
	"web_template/model/hs"

	"github.com/pkg/errors"
	"xorm.io/xorm"
)

func (s *Service) loadStudents() error {
	stus, err := s.dao.StudentList()
	if err != nil {
		log.Println("Load Students Error :", err)
		return err
	}
	s.studentList = stus
	for _, v := range stus {
		s.studentMap[v.Id] = v
	}
	return nil
}

func (s *Service) loadStudentsProc() {
	for {
		time.Sleep(time.Duration(s.c.ReloadTime) * time.Second)
		log.Println("Reload Students Data")
		s.loadStudents()
	}
}

func (s *Service) StudentAdd(req *hs.StudentAddReq) (err error) {
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

func (s *Service) StudentList() (rsp *hs.StudentListRsp, err error) {
	if s.studentList == nil {
		if err = s.loadStudents(); err != nil {
			return nil, err
		}
	}
	if len(s.studentList) == 0 {
		return nil, ecode.NothingFound
	}
	rsp = new(hs.StudentListRsp)
	for _, v := range s.studentList {
		sRsp := new(hs.StudentRsp)
		v.FormatToRsp(sRsp)
		rsp.StudentList = append(rsp.StudentList, sRsp)
	}
	return rsp, nil
}

func (s *Service) StudentById(id int) (rsp *hs.StudentRsp, err error) {
	if len(s.studentMap) == 0 {
		if err = s.loadStudents(); err != nil {
			return nil, err
		}
	}
	if v, ok := s.studentMap[id]; ok {
		rsp = new(hs.StudentRsp)
		v.FormatToRsp(rsp)
		return rsp, nil
	}
	return nil, ecode.NothingFound
}
