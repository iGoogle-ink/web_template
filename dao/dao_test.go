package dao

import (
	"flag"
	"log"
	"os"
	"testing"

	"web_template/conf"
)

var dao *Dao

func TestMain(m *testing.M) {
	log.Println("init Test")
	flag.Set("conf", "../cmd/web_template.json")
	flag.Parse()
	if err := conf.ParseConfig(); err != nil {
		panic(err)
	}
	dao = New(conf.Conf)
	exitStatus := m.Run()
	os.Exit(exitStatus)
}
