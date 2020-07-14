package main

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/simple_queue/utils"
	"strconv"
)

func main() {
	// 获取连接
	conn := utils.GetRabbitMQConnection()
	defer conn.Close()
	// 获取通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	//申明发送队列
	q, err := ch.QueueDeclare("work", true, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")
	body := "hello work queue"
	//	消息内容
	for i := 0; i < 100; i++ {
		err := ch.Publish("", q.Name, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body + " " + strconv.Itoa(i)),
		})
		failOnError(err, "Failed to publish a msg")
	}


}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}
