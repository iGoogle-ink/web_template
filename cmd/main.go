package main

import (
	"web_template/conf"
	"web_template/dao"
	"web_template/http"
)

//var consumer *nsq.Consumer

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
	//// init NSQ Producer
	//producer := dao.InitNSQProducer(conf.Conf.NSQ)
	//defer producer.Stop()
	//// init NSQ Consumer
	//consumer = dao.InitNSQConsumer(conf.Conf.NSQ)
	//defer consumer.Stop()
	// init HttpServer
	http.Init(conf.Conf, db, rds /*, producer*/)

	//ch := make(chan os.Signal)
	//
	//signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	//for {
	//	si := <-ch
	//	switch si {
	//	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
	//		time.Sleep(time.Second)
	//		log.Info("get a signal %s, stop the tv-admin process", si.String())
	//
	//		consumer.Stop()
	//		time.Sleep(time.Second)
	//		return
	//	case syscall.SIGHUP:
	//	default:
	//		return
	//	}
	//}
}
