package hs

import (
	"context"

	"web_template/conf"
	schoolDao "web_template/dao/hs"
	"web_template/model/hs"
	"web_template/pkg/errgroup"

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

func New(c *conf.Config, db *xorm.Engine, rds *redis.ClusterClient) (s *Service) {
	s = &Service{
		c:          c,
		dao:        schoolDao.New(c, db, rds),
		studentMap: make(map[int]*hs.Student),
	}

	// some data need init when service new
	s.initData()
	go s.loadStudentsProc()
	go s.loadTeachersProc()
	return s
}

func (s *Service) initData() {
	eg := errgroup.WithContext(context.Background())
	// goroutine load students
	eg.Go(func(ctx context.Context) error {
		if err := s.loadStudents(); err != nil {
			return err
		}
		return nil
	})
	// goroutine load teachers
	eg.Go(func(ctx context.Context) error {
		if err := s.loadTeachers(); err != nil {
			return err
		}
		return nil
	})
	if err := eg.Wait(); err != nil {
		panic(err)
	}
}
