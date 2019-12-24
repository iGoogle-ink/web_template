package dao

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

type NSQProducer struct {
	Producer *nsq.Producer
	Topic    string
}

type NSQConsumer struct {
	Topic string
}

func (n *NSQConsumer) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		return nil
	}
	fmt.Println("Producer Message:", string(m.Body))
	return nil
}

func (p *NSQProducer) Publish(body []byte) error {
	err := p.Producer.Publish(p.Topic, body)
	return err
}

func NewProducer(topic string, p *nsq.Producer) (producer *NSQProducer) {
	return &NSQProducer{
		Producer: p,
		Topic:    topic,
	}
}
