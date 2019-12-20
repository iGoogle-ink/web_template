package hs

import (
	"fmt"
	"log"
	"time"

	"web_template/ecode"
	"web_template/model/hs"
)

func (s *Service) loadTeachers() error {
	tchs, err := s.dao.TeacherList()
	if err != nil {
		log.Println("Load Teachers Error :", err)
		return err
	}
	var ids []int64
	for _, tch := range tchs {
		ids = append(ids, int64(tch.Id))
	}
	err = s.dao.AddCacheTeacherIds(ids)
	if err != nil {
		log.Printf("s.dao.AddCacheTeacherIds(%v) Error(%v)\n", ids, err)
		return err
	}
	err = s.dao.AddCacheTeachers(tchs)
	if err != nil {
		log.Println("s.dao.AddCacheTeachers Error :", err)
		return err
	}
	return nil
}

func (s *Service) loadTeachersProc() {
	for {
		time.Sleep(time.Duration(s.c.ReloadTime) * time.Second)
		log.Println("Reload Teachers Data")
		s.loadTeachers()
	}
}

func (s *Service) TeacherAdd(req *hs.TeacherAddReq) (err error) {
	tch := req.FormatToTeacher()
	_, err = s.dao.TeacherInsert(tch)
	if err != nil {
		fmt.Printf("s.dao.TeacherInsert(%v),err:%v.\n", *tch, err)
	}
	// todo something
	return err
}

func (s *Service) TeacherList() (rsp *hs.TeacherListRsp, err error) {
	// 获取 Teacher id
	ids, err := s.dao.CacheTeacherIds()
	if err != nil || len(ids) == 0 {
		fmt.Println("数据回源")
		tchs, err := s.dao.TeacherList()
		if err != nil {
			log.Println("s.dao.TeacherList:", err)
			return nil, err
		}
		if len(tchs) == 0 {
			return nil, ecode.NothingFound
		}
		for _, tch := range tchs {
			ids = append(ids, int64(tch.Id))
		}
		s.dao.AddCacheTeacherIds(ids)
	}
	// 根据id 获取teacher信息
	tchs, err := s.dao.CacheTeachers(ids)
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
