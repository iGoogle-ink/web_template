package main

import (
	"web_template/conf"
	"web_template/dao"
	"web_template/http"
)

func main() {
	// ParseConfig
	err := conf.ParseConfig()
	if err != nil {
		panic(err)
	}
	// init Db
	db := dao.InitDB(conf.Conf.DB)
	defer db.Close()
	// init Redis
	rds := dao.InitRedis(conf.Conf.Redis)
	defer rds.Close()
	// init HttpServer
	http.Init(conf.Conf, db, rds)
}
