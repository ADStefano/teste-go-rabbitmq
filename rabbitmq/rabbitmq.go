package rabbitmq

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func Consume(ch *amqp.Channel, out chan amqp.Delivery) {
	msgs, erro := ch.Consume(
		"Teste",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if erro != nil {
		log.Fatal(erro)
	}
	for msg := range msgs {
		out <- msg
		msg.Ack(false)
	}

}

func Send(queueName string, body string, ch *amqp.Channel) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	erro := ch.PublishWithContext(ctx,
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if erro != nil {
		log.Fatal(erro)
	}

	log.Printf(" Enviou: %s\n", body)
}

func DeclareQueue(queueName string, ch *amqp.Channel) error {
	_, erro := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if erro != nil {
		return erro
	}

	return nil
}
