package school

import (
	"web_template/conf"
	"web_template/model/school"

	"github.com/go-redis/redis/v7"
	schoolDao "web_template/dao/school"
	"xorm.io/xorm"
)

type Service struct {
	c   *conf.Config
	dao *schoolDao.Dao

	// 数据初始化，1W 条以上数据，存Redis，数据里小，可以存内存
	studentList []*school.Student
	studentMap  map[int]*school.Student // key:id,value:name
}

func New(c *conf.Config, db *xorm.Engine, rds *redis.Client) (s *Service) {
	s = &Service{
		c:          c,
		dao:        schoolDao.New(c, db, rds),
		studentMap: make(map[int]*school.Student),
	}

	// some data need init when service new
	s.loadStudent()
	go s.loadStudentProc()
	return s
}
