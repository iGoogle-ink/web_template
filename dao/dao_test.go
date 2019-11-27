package dao

import (
	"flag"
	"os"
	"testing"

	"web_template/conf"
)

var dao *Dao

func TestMain(m *testing.M) {
	flag.Set("conf", "../cmd/web_template.json")
	flag.Parse()
	if err := conf.ParseConfig(); err != nil {
		panic(err)
	}
	dao = New(conf.Conf)
	os.Exit(m.Run())
}
