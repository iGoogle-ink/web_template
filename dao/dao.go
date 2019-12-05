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

func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
	if d.Redis != nil {
		d.Redis.Close()
	}
}

func (d *Dao) Transact(transactHandler func(tx *xorm.Session) error) (err error) {
	session := d.DB.NewSession()
	if err = session.Begin(); err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = session.Rollback()
			log.Println("Panic In Transact:", p)
			return
		}
		if err != nil {
			_ = session.Rollback()
			return
		}
		err = session.Commit()
	}()

	err = transactHandler(session)
	return err
}
