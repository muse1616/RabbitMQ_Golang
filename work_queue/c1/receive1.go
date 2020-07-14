package main

import (
	"log"
	"rabbitmq/simple_queue/utils"
	"time"
)

func main() {
	// 获取rabbitMQ连接
	conn := utils.GetRabbitMQConnection()
	// 关闭连接
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("work", true, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	//注册消费者
	//  关闭自动确认 autAck:false 同时一次只给一个消息
	err = ch.Qos(1, 0, false)
	failOnError(err, "Failed to set qos")
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	//通道 异步读取 开启通道是为了阻塞
	forever := make(chan []byte)

	go func() {
		for d := range msgs {
			// 延迟1秒
			t := time.Duration(1)
			time.Sleep(t * time.Second)
			log.Printf("消费者1:Received a message:%s\n", d.Body)
			//是否多个消息同时确认 false 即可
			_ = d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}
