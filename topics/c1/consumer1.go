package main

import (
	"log"
	"rabbitmq/simple_queue/utils"
)

func main() {
	conn := utils.GetRabbitMQConnection()
	defer conn.Close()
	ch, _ := conn.Channel()
	defer ch.Close()
	_ = ch.ExchangeDeclare("topics", "topic", true, false, false, false, nil)
	q, _ := ch.QueueDeclare("", false, false, true, false, nil)
	// 通配符绑定
	_ = ch.QueueBind(q.Name, "user.*", "topics", false, nil)

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
