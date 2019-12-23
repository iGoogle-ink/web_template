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

func InitRedis(c *conf.Redis) (r *redis.ClusterClient) {
	r = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    c.Addr,
		Password: c.Password,
	})
	_, err := r.Ping().Result()
	if err != nil {
		panic(err)
	}
	return r
}
