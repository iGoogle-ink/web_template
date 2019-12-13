package auth

import (
	"flag"
	"os"
	"testing"

	"web_template/conf"
	"web_template/dao"
)

var d *Dao

func TestMain(m *testing.M) {
	flag.Set("env", "test")
	flag.Set("conf", "../../cmd/web_template.json")
	flag.Parse()
	if err := conf.ParseConfig(); err != nil {
		panic(err)
	}
	rds := dao.InitRedis(conf.Conf.Redis)
	d = New(conf.Conf, rds)
	os.Exit(m.Run())
}
