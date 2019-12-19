package hs

import (
	"fmt"
	"log"

	"web_template/ecode"
	"web_template/model/hs"
)

func (s *Service) TeacherAdd(req *hs.TeacherAddReq) (err error) {
	tch := req.FormatToTeacher()
	_, err = s.dao.TeacherInsert(tch)
	if err != nil {
		fmt.Printf("s.dao.TeacherInsert(%v),err:%v.\n", *tch, err)
	}
	// todo something
	return err
}

func (s *Service) TeacherList(start, end int64) (rsp *hs.TeacherListRsp, err error) {
	var tchs []*hs.Teacher
	tchs, err = s.dao.CacheTeacher(start, end)
	if err != nil {
		return nil, err
	}
	if len(tchs) == 0 {
		fmt.Println("数据回源")
		tchs, err = s.dao.TeacherList()
		if err != nil {
			log.Println("s.dao.TeacherList:", err)
			return nil, err
		}
		if len(tchs) == 0 {
			return nil, ecode.NothingFound
		}
		cacheErr := s.dao.AddCacheTeacher(tchs)
		if cacheErr != nil {
			log.Println("缓存添加失败")
		}
	}
	rsp = new(hs.TeacherListRsp)
	for _, v := range tchs {
		t := new(hs.TeacherRsp)
		v.FormatToRsp(t)
		rsp.TeacherList = append(rsp.TeacherList, t)
	}
	// todo something
	return rsp, nil
}

func (s *Service) TeacherById(id int) (*hs.Teacher, error) {
	tch, err := s.dao.TeacherById(id)
	if err != nil {
		return nil, err
	}
	return tch, nil
}
