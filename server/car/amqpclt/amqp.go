package amqpclt

import (
	"context"
	carpb "coolcar/car/api/gen/v1"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	ch       *amqp.Channel
	exchange string
}

func NewPublisher(conn *amqp.Connection, exchange string) (*Publisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("cannot allocate channel: %v", err)
	}

	err = ch.ExchangeDeclare(
		exchange,
		"fanout",
		true,  // durable
		false, // autoDelete
		false, // internal
		false, // noWait
		nil,   //args
	)
	if err != nil {
		return nil, fmt.Errorf("cannot declare exchange: %v", err)
	}

	return &Publisher{
		ch:       ch,
		exchange: exchange,
	}, nil
}

func (p *Publisher) Publish(c context.Context, car *carpb.CarEntity) error {
	b, err := json.Marshal(car)
	if err != nil {
		return fmt.Errorf("cannot marshal: %v", err)
	}

	return p.ch.PublishWithContext(
		c,
		p.exchange,
		"",    // key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		},
	)
}
