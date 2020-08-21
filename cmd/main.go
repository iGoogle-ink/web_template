package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"gotil/conf"
	"gotil/xlog"
	"web_template/config"
	"web_template/router"
	"web_template/service"
)

func main() {
	// Parse Config
	err := conf.ParseYaml(config.Conf)
	if err != nil {
		panic(err)
	}
	// New Service
	svc := service.New(config.Conf)
	// Start Web Server
	router.Init(config.Conf, svc)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			xlog.Warningf("%s: get a signal %s, stop the process", config.Conf.Name, si.String())
			svc.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
