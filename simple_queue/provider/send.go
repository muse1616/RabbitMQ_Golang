package main

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/simple_queue/utils"
)

/**
	消息队列 生产者
 */
func main() {
	// 获取rabbitMq连接
	conn := utils.GetRabbitMQConnection()
	// 关闭连接
	defer conn.Close()
	// 通过连接 创建一个通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	//关闭通道
	defer ch.Close()
	// 通过通道 声明一个队列
	//参数1：队列名称 队列不存在时 会自动创建
	//参数2: 是否持久化 即关闭队列后是否存储到磁盘 下次打开仍然存在 只保证队列持久化 队列中未消费的消息仍然会消失 需要在发布消息时明确定义
	//参数3： 使用完毕后是否自动关闭
	//参数4: 该队列是否是该连接独占
	//参数5: 是否非阻塞 等待队列
	//参数6: 其余可选参数 nil即可
	q, err := ch.QueueDeclare("hello", true, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	//	消息内容
	body := "hello world"
	err = ch.Publish(
		// 交换机 简单队列无需交换机
		"",
		//队列名称
		q.Name,
		// 默认为false即可
		false,
		false,
		// 消息发布
		amqp.Publishing{
			// 消息持久化 *
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}

// 检查每个 amqp 调用返回值
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}
