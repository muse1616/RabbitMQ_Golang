package main

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/simple_queue/utils"
	"strconv"
)

func main() {
	conn := utils.GetRabbitMQConnection()
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	q, _ := ch.QueueDeclare(
		"rpc_queue",
		false,
		false,
		false,
		false,
		nil,
	)

	// 一次只拿一条消息
	_ = ch.Qos(
		1,
		0,
		false,
	)

	//消费信息
	msgs, _ := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)

	//协程从msgs chan中读取数据
	go func() {
		for d := range msgs {
			//转化为数字
			n, _ := strconv.Atoi(string(d.Body))

			//开辟一个协程处理
			d := d
			go func() {

				// 此处待修改 需要fork子进程去完成任务

				log.Printf(" [.] fib(%d)", n)
				response := fib(n)
				////模拟处理10秒
				//t := time.Duration(10)
				//time.Sleep(t * time.Second)
				// 回调队列
				_ = ch.Publish(
					"",        // exchange
					d.ReplyTo, // routing key
					false,     // mandatory
					false,     // immediate
					amqp.Publishing{
						ContentType:   "text/plain",
						CorrelationId: d.CorrelationId,
						Body:          []byte(strconv.Itoa(response)),
					})
				_ = d.Ack(false)
			}()
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}

// 计算斐波那契
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
