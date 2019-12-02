package service

import "web_template/model"

func (s *Service) TeacherAdd(req *model.TeacherAddReq) (err error) {
	tch := req.FormatToTeacher()
	_, err = s.dao.TeacherInsert(tch)
	// todo something
	return err
}

func (s *Service) TeacherList() (rsp *model.TeacherListRsp, err error) {
	list, err := s.dao.TeacherList()
	if err != nil {
		return nil, err
	}
	rsp = new(model.TeacherListRsp)
	for _, v := range list {
		t := new(model.TeacherRsp)
		v.FormatToRsp(t)
		rsp.TeacherList = append(rsp.TeacherList, t)
	}
	// todo something
	return rsp, nil
}
