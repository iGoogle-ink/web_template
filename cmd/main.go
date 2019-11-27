package main

import (
	"web_template/conf"
	"web_template/http"
	"web_template/service"
)

func main() {
	// ParseConfig
	err := conf.ParseConfig()
	if err != nil {
		panic(err)
	}
	// init Service
	svr := initService(conf.Conf)
	// init HttpServer
	http.Init(conf.Conf, svr)
}

func initService(c *conf.Config) (svr *http.Server) {
	svr = &http.Server{
		Service: service.New(c),
	}
	return svr
}
