package dao

import (
	"log"

	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"web_template/conf"
	"xorm.io/xorm"
)

const (
	_TableBinding = "binding"
	_TableStudent = "student"
	_TableTeacher = "teacher"
)

type Dao struct {
	c     *conf.Config
	DB    *xorm.Engine
	Redis *redis.Client
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:     c,
		DB:    initDB(c.DB),
		Redis: initRedis(c.Redis),
	}
	return d
}

func initDB(c *conf.DB) (db *xorm.Engine) {
	db, err := xorm.NewEngine(c.Driver, c.Dsn)
	if err != nil {
		panic(err)
	}
	db.ShowSQL(c.ShowSQL)
	db.SetMaxIdleConns(c.Idle)
	return db
}

func initRedis(c *conf.Redis) (r *redis.Client) {
	r = redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.Db,
	})
	_, err := r.Ping().Result()
	if err != nil {
		log.Println("InitRedis Error:", err)
		panic(err)
	}
	return r
}

func (d *Dao) EndTransact(tx *xorm.Session, err error) error {
	if err != nil {
		if e := tx.Rollback(); e != nil {
			return e
		}
		return err
	}
	return tx.Commit()
}

//func (d *Dao) Transact(tx *xorm.Session, transactHandler func(tx *xorm.Session) error) error {
//	if err := transactHandler; err != nil {
//		return tx.Rollback()
//	}
//	return tx.Commit()
//}
