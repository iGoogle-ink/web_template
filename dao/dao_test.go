package dao

import (
	"context"
	"flag"
	"os"
	"testing"

	"gotil/conf"
	"web_template/config"
)

var (
	d   *Dao
	ctx = context.Background()
)

func TestMain(m *testing.M) {
	if os.Getenv("UT_LOCAL_TEST") != "" {
		os.Setenv("RUNTIME_ENV", "local")
	}
	flag.Set("conf", "../cmd/config.yaml")
	if err := conf.ParseYaml(config.Conf); err != nil {
		panic(err)
	}
	d = New(config.Conf)
	os.Exit(m.Run())
}
