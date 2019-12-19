package hs

import (
	"web_template/conf"
	schoolDao "web_template/dao/hs"
	"web_template/model/hs"

	"github.com/go-redis/redis/v7"
	"xorm.io/xorm"
)

type Service struct {
	c   *conf.Config
	dao *schoolDao.Dao

	// 数据初始化，1W 条以上数据，存Redis，数据里小，可以存内存
	studentList []*hs.Student
	studentMap  map[int]*hs.Student // key:id,value:name
}

func New(c *conf.Config, db *xorm.Engine, rds *redis.Client) (s *Service) {
	s = &Service{
		c:          c,
		dao:        schoolDao.New(c, db, rds),
		studentMap: make(map[int]*hs.Student),
	}

	// some data need init when service new
	s.loadStudent()
	go s.loadStudentProc()
	return s
}
