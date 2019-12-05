package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"web_template/conf"
	"web_template/http"
	"web_template/service"
)

var svr *http.Server

func main() {
	// ParseConfig
	err := conf.ParseConfig()
	if err != nil {
		panic(err)
	}
	// init Service
	svr = initService(conf.Conf)
	// init HttpServer
	http.Init(conf.Conf, svr)
	signalHandler()
}

func initService(c *conf.Config) (svr *http.Server) {
	svr = &http.Server{
		Service: service.New(c),
	}
	return svr
}

func signalHandler() {
	var ch = make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Printf("get a signal %s, stop the consume process.\n", si.String())
			svr.Close()
			// todo: some other Close

			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
