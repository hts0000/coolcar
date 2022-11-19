package amqpclt

import (
	"context"
	carpb "coolcar/car/api/gen/v1"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
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

	err = declareExchange(ch, exchange)
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

type Subscriber struct {
	conn     *amqp.Connection
	logger   *zap.Logger
	exchange string
}

func NewSubscriber(conn *amqp.Connection, exchange string, logger *zap.Logger) (*Subscriber, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("cannot allocate channel: %v", err)
	}
	defer ch.Close()

	err = declareExchange(ch, exchange)
	if err != nil {
		return nil, fmt.Errorf("cannot declare exchange: %v", err)
	}

	return &Subscriber{
		conn:     conn,
		logger:   logger,
		exchange: exchange,
	}, nil
}

func (s *Subscriber) SubscribeRaw(c context.Context) (<-chan amqp.Delivery, func(), error) {
	ch, err := s.conn.Channel()
	if err != nil {
		return nil, func() {}, fmt.Errorf("cannot allocate channel: %v", err)
	}

	closeCh := func() {
		err := ch.Close()
		if err != nil {
			s.logger.Error("cannot close channel", zap.Error(err))
		}
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, //durable
		true,  // autoDelete
		false, // exclusive
		false, // noWait
		nil,   //agrs
	)
	if err != nil {
		return nil, closeCh, fmt.Errorf("cannot declare queue: %v", err)
	}

	cleanUp := func() {
		_, err := ch.QueueDelete(
			q.Name, // queue name
			false,  // ifUnused
			false,  // ifEmpty
			false,  // noWait
		)
		if err != nil {
			s.logger.Error("cannot delete queue", zap.String("name", q.Name), zap.Error(err))
		}

		closeCh()
	}

	err = ch.QueueBind(
		q.Name,     // queue name
		"",         //key
		s.exchange, // exchange name
		false,      //noWait
		nil,        //args
	)
	if err != nil {
		return nil, cleanUp, fmt.Errorf("cannot bind queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer
		true,   // autoAck
		false,  // exclusive
		false,  // noLocal
		false,  // noWait
		nil,    // agrs
	)
	if err != nil {
		return nil, cleanUp, fmt.Errorf("cannot comsume queue: %v", err)
	}

	return msgs, cleanUp, nil
}

func (s *Subscriber) Subscribe(c context.Context) (chan *carpb.CarEntity, func(), error) {
	msgCh, cleanUp, err := s.SubscribeRaw(c)
	if err != nil {
		return nil, cleanUp, err
	}

	carCh := make(chan *carpb.CarEntity)
	go func() {
		for msg := range msgCh {
			var car carpb.CarEntity
			err := json.Unmarshal(msg.Body, &car)
			if err != nil {
				s.logger.Error("cannot unmarshal", zap.Error(err))
			}
			carCh <- &car
		}
		close(carCh)
	}()

	return carCh, cleanUp, nil
}

func declareExchange(ch *amqp.Channel, exchange string) error {
	return ch.ExchangeDeclare(
		exchange,
		"fanout",
		true,  // durable
		false, // autoDelete
		false, // internal
		false, // noWait
		nil,   //args
	)
}
