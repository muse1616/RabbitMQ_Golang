package main

import (
	"log"
	"rabbitmq/simple_queue/utils"
)

func main() {
	conn := utils.GetRabbitMQConnection()
	ch, _ := conn.Channel()
	_ = ch.ExchangeDeclare("logs_direct", "direct", true, false, false, false, nil)
	//	tmp Queue
	q, _ := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	//	绑定队列
	_ = ch.QueueBind(
		q.Name,
		"info2",
		"logs_direct",
		false,
		nil,
	)

	msgs, _ := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("[x] %s", d.Body)
		}
	}()
	log.Printf("1 [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
