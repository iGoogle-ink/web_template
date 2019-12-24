package pay

import (
	"web_template/conf"
	"web_template/dao"
	"web_template/dao/pay"

	"github.com/go-redis/redis/v7"
	"github.com/nsqio/go-nsq"
	"xorm.io/xorm"
)

type Service struct {
	c        *conf.Config
	dao      *pay.Dao
	producer *dao.NSQProducer
}

func New(c *conf.Config, db *xorm.Engine, rds *redis.ClusterClient, producer *nsq.Producer) (s *Service) {
	s = &Service{
		c:        c,
		dao:      pay.New(c, db, rds),
		producer: dao.NewProducer(c.NSQ.Topic, producer),
	}
	return s
}
