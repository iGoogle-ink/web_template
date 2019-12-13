package pay

import (
	"github.com/go-redis/redis/v7"
	"web_template/conf"
	"web_template/dao/pay"
	"xorm.io/xorm"
)

type Service struct {
	c   *conf.Config
	dao *pay.Dao
}

func New(c *conf.Config, db *xorm.Engine, rds *redis.Client) (s *Service) {
	s = &Service{
		c:   c,
		dao: pay.New(c, db, rds),
	}
	return s
}
