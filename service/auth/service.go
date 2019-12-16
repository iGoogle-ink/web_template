package auth

import (
	"web_template/conf"
	"web_template/dao/auth"

	"github.com/go-redis/redis/v7"
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
