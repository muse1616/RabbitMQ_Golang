package main

import (
	"log"
	"rabbitmq/simple_queue/utils"
)

/**
消息队列 消费者
*/
func main() {
	// 获取rabbitMQ连接
	conn := utils.GetRabbitMQConnection()
	// 关闭连接
	defer conn.Close()

	// 通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 队列 注意参数需要和生产者完全相同
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	// 注册消费者
	// autoAck 自动确认
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	//通道 异步读取 开启通道是为了阻塞
	forever := make(chan []byte)

	// 协程读取
	go func() {
		for d := range msgs {
			log.Printf("Received a message:%s\n", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// 无限阻塞
	<-forever
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}
