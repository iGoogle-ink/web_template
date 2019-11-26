package main

import (
	"web_template/conf"
	"web_template/http"
)

func main() {
	err := conf.ParseConfig()
	if err != nil {
		panic(err)
	}
	http.Init(conf.Conf)
}
