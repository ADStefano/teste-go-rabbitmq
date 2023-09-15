package main

import (
	"fmt"
	"log"
	"teste-go-rabbit/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	channel, erro := rabbitmq.OpenChannel()
	if erro != nil {
		log.Fatal("Falha ao abrir conexão com o rabbit")
	}

	defer channel.Close()

	erro = rabbitmq.DeclareQueue("Teste", channel)
	if erro != nil {
		log.Fatal(erro)
	}

	erro = rabbitmq.DeclareQueue("Gorlami", channel)
	if erro != nil {
		log.Fatal(erro)
	}

	msgRabbitMqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(channel, msgRabbitMqChannel)

	fmt.Printf("Canal: %d", msgRabbitMqChannel)

	for msg := range msgRabbitMqChannel{
		fmt.Printf("Mensagem: %s", string(msg.Body))
		for i := 0; i < 700000; i++ {
			go rabbitmq.Send("Gorlami", "teste", channel)
		}
	}

	forever := make(chan int)
	<-forever

}
