package service

import (
	"log"

	"web_template/model"
)

func (s *Service) TeacherAdd(req *model.TeacherAddReq) (err error) {
	tch := req.FormatToTeacher()
	_, err = s.dao.TeacherInsert(tch)
	// todo something
	return err
}

func (s *Service) TeacherList() (rsp *model.TeacherListRsp, err error) {
	tchs, err := s.dao.CacheTeacher(0, 100)
	if err != nil || len(tchs) == 0 {
		log.Println("s.dao.CacheTeacher: err or empty:", err)
		err = nil
		tchs, err = s.dao.TeacherList()
		if err != nil {
			return nil, err
		}
		s.dao.AddCacheTeacher(tchs)
	}
	rsp = new(model.TeacherListRsp)
	for _, v := range tchs {
		t := new(model.TeacherRsp)
		v.FormatToRsp(t)
		rsp.TeacherList = append(rsp.TeacherList, t)
	}
	// todo something
	return rsp, nil
}
