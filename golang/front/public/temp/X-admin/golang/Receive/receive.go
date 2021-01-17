package main

import (
	"Send"
	"github.com/streadway/amqp"
	"log"
	"seckill2"
)

func main(){
	// 连接RabbitMQ服务器
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Send.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// 创建一个channel
	ch, err := conn.Channel()
	Send.FailOnError(err, "Failed to open a channel")
	defer ch.Close()
	// 监听队列
	q0, err  := ch.QueueDeclare(
		"username",			// 队列名称
		false,			// 是否持久化
		false,		// 是否自动删除
		false,			// 是否独立
		false,nil,
	)
	Send.FailOnError(err, "Failed to declare a queue")
	q1, err  := ch.QueueDeclare(
		"name",			// 队列名称
		false,			// 是否持久化
		false,		// 是否自动删除
		false,			// 是否独立
		false,nil,
	)
	Send.FailOnError(err, "Failed to declare a queue")
	q2, err  := ch.QueueDeclare(
		"max_nums",			// 队列名称
		false,			// 是否持久化
		false,		// 是否自动删除
		false,			// 是否独立
		false,nil,
	)
	Send.FailOnError(err, "Failed to declare a queue")
	// 消费队列
	msgs0, err := ch.Consume(
		q0.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	Send.FailOnError(err, "Failed to register a consumer")
	msgs1, err := ch.Consume(
		q1.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	Send.FailOnError(err, "Failed to register a consumer")
	msgs2, err := ch.Consume(
		q2.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	Send.FailOnError(err, "Failed to register a consumer")
	// 申明一个goroutine,一遍程序始终监听
	forever := make(chan bool)

	go func() {
		for d := range msgs0{
			for d1 := range msgs1 {
				for d2 := range msgs2 {
					var str string =string(d2.Body)
					var a int=int(str[0]-'0')
					log.Printf("Received a message: %s,%s,%d",d.Body, d1.Body,a)
					seckill2.Deal(string (d.Body),string(d1.Body),a)
					break
				}
			}
		}

	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// 帮助函数检测每一个amqp调用
