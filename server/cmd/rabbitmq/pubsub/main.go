package main

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const exchange = "go_ex"

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@hts0000.top:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	go subscribe(exchange, conn)
	go subscribe(exchange, conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for i := 0; ; i++ {
		err := ch.PublishWithContext(
			ctx,
			exchange,
			"",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(fmt.Sprintf("message: %d", i)),
			},
		)
		if err != nil {
			panic(err)
		}
		time.Sleep(250 * time.Millisecond)
	}
}

func subscribe(ex string, conn *amqp.Connection) {
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	defer ch.QueueDelete(
		q.Name,
		false,
		false,
		false,
	)

	err = ch.QueueBind(
		q.Name,
		"",
		ex,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	consume("c", ch, q.Name)
}

func consume(consumer string, ch *amqp.Channel, q string) {
	msgs, err := ch.Consume(
		q,
		consumer,
		true,  // autoAsk
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,   // agrs
	)
	if err != nil {
		panic(err)
	}

	for msg := range msgs {
		fmt.Printf("%s: %s\n", consumer, msg.Body)
	}
}
