package main

import (
	"log"
	"rabbitmq/simple_queue/utils"
)

func main() {
	//	获取连接
	conn := utils.GetRabbitMQConnection()
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "channel")
	defer ch.Close()
	//交换机
	err = ch.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "exchange declare")
	//	临时队列
	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "tmp queue")
	//	绑定队列
	err = ch.QueueBind(
		q.Name,
		"",
		"logs",
		false,
		nil,
	)
	failOnError(err, "bind queue")
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("[x] %s", d.Body)
		}
	}()
	log.Printf(" 2 [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}
