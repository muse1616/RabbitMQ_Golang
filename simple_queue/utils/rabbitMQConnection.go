package utils

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//rabbitMQ连接工具类
func GetRabbitMQConnection() *amqp.Connection {
	// 连接url 账号 密码 ip port 注意url最后加上 vhost 有反斜线需要保留反斜线
	url := fmt.Sprintf("amqp://%s:%s@%s:%d//test", "admin", "admin", "192.168.88.130", 5672)
	// 获取连接
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ Server")
	return conn
}
//错误打印函数
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}
