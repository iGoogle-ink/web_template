package hs

import (
	"log"

	"web_template/conf"

	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

const (
	_TableBinding   = "binding"
	_TableHsStudent = "hs_student"
	_TableHsTeacher = "hs_teacher"
)

type Dao struct {
	c     *conf.Config
	DB    *xorm.Engine
	Redis *redis.ClusterClient
}

func New(c *conf.Config, db *xorm.Engine, rds *redis.ClusterClient) (d *Dao) {
	d = &Dao{
		c:     c,
		DB:    db,
		Redis: rds,
	}
	return d
}

func (d *Dao) Transact(transactHandler func(tx *xorm.Session) error) (err error) {
	session := d.DB.NewSession()
	if err = session.Begin(); err != nil {
		return err
	}

	defer func() {
		defer session.Close()
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
