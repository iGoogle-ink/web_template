package auth

import (
	"web_template/conf"

	"github.com/go-redis/redis/v7"
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
