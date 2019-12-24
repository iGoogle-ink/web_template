package dao

import (
	"fmt"
	"time"

	"web_template/conf"

	"github.com/go-redis/redis/v7"
	"github.com/nsqio/go-nsq"
	"xorm.io/xorm"
)

func InitDB(c *conf.DB) (db *xorm.Engine) {
	db, err := xorm.NewEngine(c.Driver, c.Dsn)
	if err != nil {
		panic(err)
	}
	db.ShowSQL(c.ShowSQL)
	db.SetMaxIdleConns(c.Idle)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func InitRedis(c *conf.Redis) (r *redis.ClusterClient) {
	r = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    c.Addr,
		Password: c.Password,
	})
	_, err := r.Ping().Result()
	if err != nil {
		panic(err)
	}
	return r
}

func InitNSQConsumer(c *conf.NSQ) (consumer *nsq.Consumer) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 10 * time.Second
	consumer, err := nsq.NewConsumer(c.Topic, c.Channel, config)
	if err != nil {
		panic(err)
	}
	handler := &NSQConsumer{Topic: c.Topic}
	consumer.AddHandler(handler)
	fmt.Println("consumer:", c.ConsumerAddr)
	err = consumer.ConnectToNSQLookupd(c.ConsumerAddr)
	if err != nil {
		panic(err)
	}
	return consumer
}

func InitNSQProducer(c *conf.NSQ) (producer *nsq.Producer) {
	config := nsq.NewConfig()
	fmt.Println("producer:", c.ProducerAddr)
	producer, err := nsq.NewProducer(c.ProducerAddr, config)
	if err != nil {
		panic(err)
	}
	return producer
}
