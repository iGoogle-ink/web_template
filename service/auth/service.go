package auth

import (
	"github.com/go-redis/redis/v7"
	"web_template/conf"
	"web_template/dao/auth"
)

type Service struct {
	c   *conf.Config
	dao *auth.Dao
}

func New(c *conf.Config, rds *redis.Client) (s *Service) {
	s = &Service{
		c:   c,
		dao: auth.New(c, rds),
	}
	return s
}
