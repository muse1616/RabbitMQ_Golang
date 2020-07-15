package main

import (
	"log"
	"rabbitmq/simple_queue/utils"
	"strconv"
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
	// 可以绑定多个队列 接受多个routineKey消息
	for i := 1; i <= 2; i++ {
		_ = ch.QueueBind(
			q.Name,
			"info"+strconv.Itoa(i),
			"logs_direct",
			false,
			nil,
		)
	}

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
