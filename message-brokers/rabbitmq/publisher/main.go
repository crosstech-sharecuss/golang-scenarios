package main

import (
	"github.com/streadway/amqp"
)

func main() {
	// RabbitMQ server
	amqpServerURL := "amqp://rabbitmq-address"

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// With the instance and declare Queues that we can
	// publish and subscribe to.
	_, err = channelRabbitMQ.QueueDeclare(
		"AgentTestQueue", // queue name
		true,             // durable
		false,            // auto delete
		false,            // exclusive
		false,            // no wait
		nil,              // arguments
	)
	if err != nil {
		panic(err)
	}

	// Create a message to publish.
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello World 2!"),
	}

	// Attempt to publish a message to the queue.
	if err := channelRabbitMQ.Publish(
		"AgentTestExchange", // exchange
		"AgentTestQueue",    // routing-key /or/ queue name if no exchange provided
		false,               // mandatory
		false,               // immediate
		message,             // message to publish
	); err != nil {
		panic(err)
	}
}
