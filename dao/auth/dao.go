package auth

import (
	"github.com/go-redis/redis/v7"
	"web_template/conf"
)

type Dao struct {
	c     *conf.Config
	Redis *redis.Client
}

func New(c *conf.Config, rds *redis.Client) (d *Dao) {
	d = &Dao{
		c:     c,
		Redis: rds,
	}
	return d
}
