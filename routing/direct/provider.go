package main

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/simple_queue/utils"
	"strconv"
)

func main() {
	// 连接
	conn := utils.GetRabbitMQConnection()
	defer conn.Close()
	//	通道
	ch, err := conn.Channel()
	failOnError(err, "channel")
	defer ch.Close()
	//	交换机 路由模式direct
	err = ch.ExchangeDeclare("logs_direct", "direct", true, false, false, false, nil)
	//	发送消息
	//	制定路由key
	routineKey := "info"
	body := "direct: "
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			err = ch.Publish(
				"logs_direct",
				routineKey+"1",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body + strconv.Itoa(i)),
				},
			)
			failOnError(err, "publish")
		} else {
			err = ch.Publish(
				"logs_direct",
				routineKey+"2",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body + strconv.Itoa(i)),
				},
			)
			failOnError(err, "publish")
		}

	}

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}
