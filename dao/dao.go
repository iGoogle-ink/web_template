package dao

import (
	"gotil/orm"
	"web_template/config"

	"github.com/jinzhu/gorm"
)

type Dao struct {
	MySQL *gorm.DB
	redis *orm.Redis
	c     *config.Config
}

func New(c *config.Config) (d *Dao) {
	d = &Dao{
		MySQL: orm.InitGorm(c.MySQL),
		redis: orm.InitRedis(c.Redis),
		c:     c,
	}
	return d
}

func (d *Dao) Close() {
	if d.MySQL != nil {
		d.MySQL.Close()
	}
}
