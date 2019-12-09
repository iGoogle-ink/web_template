package service

import (
	"fmt"

	"web_template/ecode"
	"web_template/model"
)

func (s *Service) TeacherAdd(req *model.TeacherAddReq) (err error) {
	tch := req.FormatToTeacher()
	_, err = s.dao.TeacherInsert(tch)
	if err != nil {
		fmt.Printf("s.dao.TeacherInsert(%v),err:%v.\n", *tch, err)
	}
	// todo something
	return err
}

func (s *Service) TeacherList(start, end int64) (rsp *model.TeacherListRsp, err error) {
	var tchs []*model.Teacher
	tchs, err = s.dao.CacheTeacher(start, end)
	if err != nil {
		return nil, err
	}
	if len(tchs) == 0 {
		fmt.Println("数据回源")
		tchs, err = s.dao.TeacherList()
		if err != nil {
			return nil, err
		}
		if len(tchs) == 0 {
			return nil, ecode.NothingFound
		}
		cacheErr := s.dao.AddCacheTeacher(tchs)
		if cacheErr != nil {
			fmt.Println("缓存添加失败")
		}
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
