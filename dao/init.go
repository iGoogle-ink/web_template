package dao

import (
	"web_template/conf"

	"github.com/go-redis/redis/v7"
	"xorm.io/xorm"
)

func InitDB(c *conf.DB) (db *xorm.Engine) {
	db, err := xorm.NewEngine(c.Driver, c.Dsn)
	if err != nil {
		panic(err)
	}
	db.ShowSQL(c.ShowSQL)
	db.SetMaxIdleConns(c.Idle)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func InitRedis(c *conf.Redis) (r *redis.Client) {
	r = redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.Db,
	})
	_, err := r.Ping().Result()
	if err != nil {
		panic(err)
	}
	return r
}
