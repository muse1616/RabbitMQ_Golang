package main

import (
	"github.com/streadway/amqp"
	"rabbitmq/simple_queue/utils"
)

func main() {
	//	连接
	conn := utils.GetRabbitMQConnection()
	defer conn.Close()
	ch, _ := conn.Channel()
	defer ch.Close()
	_ = ch.ExchangeDeclare("topics", "topic", true, false, false, false, nil)
	//	发布消息
	routineKey := "user.save"
	_ = ch.Publish("topics", routineKey, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(routineKey)})

}
