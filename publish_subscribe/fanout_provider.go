package main

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/simple_queue/utils"
)

func main() {
	conn := utils.GetRabbitMQConnection()
	defer conn.Close()
	ch, err := conn.Channel()

	defer ch.Close()

	if ch != nil {
		defer ch.Close()
	}
	failOnError(err, "channel")
	// 参数1:交换机名字
	// 参数2:交换机类型 广播是 fanout
	// 参数3:是否持久化
	// 参数4:是否自动删除
	// 参数5:
	// 参数6:
	// 参数7:其余参数
	err = ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)
	failOnError(err, "exchange")

	//	发送消息
	body := "hello exchange"
	err = ch.Publish(
		// 交换机名称
		"logs",
		//  key为空即可
		"",
		// 默认false即可
		false,
		false,
		amqp.Publishing{
			//消息持久化
			//DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body:        []byte(body),
		})
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}
