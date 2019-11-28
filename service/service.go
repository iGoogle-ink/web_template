package service

import (
	"web_template/conf"
	"web_template/dao"
	"web_template/model"
)

type Service struct {
	c   *conf.Config
	dao *dao.Dao

	// 数据初始化，1W 条以上数据，存Redis，数据里小，可以存内存
	studentList []*model.Student
	studentMap  map[int]*model.Student // key:id,value:name
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:          c,
		dao:        dao.New(c),
		studentMap: make(map[int]*model.Student),
	}

	// some data need init when service new
	s.loadStudent()
	go s.loadStudentProc()
	return s
}
