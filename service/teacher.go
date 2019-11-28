package service

import "web_template/model"

func (s *Service) TeacherAdd(req *model.TeacherAddReq) (err error) {
	_, err = s.dao.TeacherInsert()
	// todo something
	return err
}
