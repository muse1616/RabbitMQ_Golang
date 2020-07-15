package main

import (
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"rabbitmq/simple_queue/utils"
	"strconv"
	"time"
)

func main() {
	// 生成一个随机数
	rand.Seed(time.Now().UTC().UnixNano())
	n := randInt(10, 20)
	log.Printf(" [x] Requesting fib(%d)", n)


	//获取连接
	conn := utils.GetRabbitMQConnection()
	defer conn.Close()
	//通道
	ch, _ := conn.Channel()
	defer ch.Close()
	//简单队列
	q, _ := ch.QueueDeclare(
		"",
		false,
		false,
		//独占队列
		true,
		false,
		nil,
	)
	//发布者 回调队列 消费
	msgs, _ := ch.Consume(
		q.Name, // queue
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	//corrID 随机生成
	corrId := randomString(32)

	//publish消息
	_ = ch.Publish(
		"",
		"rpc_queue",
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			// 绑定id
			CorrelationId: corrId,
			//回调队列
			ReplyTo:       q.Name,
			Body:          []byte(strconv.Itoa(n)),
		})

	for d := range msgs {
		if corrId == d.CorrelationId {
			res, _ := strconv.Atoi(string(d.Body))
			log.Printf(" [.] Got %d", res)
			break
		}
	}
}

// 生成随机长度为l的字符串 用来随机生成corrID
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

// 生成[max,min]中的随机数
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
