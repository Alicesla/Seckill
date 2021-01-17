package Send
import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)
func S(a string,b string,c string){
	// 连接RabbitMQ服务器
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// 创建一个channel
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个队列
	q0, err  := ch.QueueDeclare(
		"username",			// 队列名称
		false,			// 是否持久化
		false,		// 是否自动删除
		false,			// 是否独立
		false,nil,
	)
	FailOnError(err, "Failed to declare a queue")
	q1, err  := ch.QueueDeclare(
		"name",			// 队列名称
		false,			// 是否持久化
		false,		// 是否自动删除
		false,			// 是否独立
		false,nil,
	)
	FailOnError(err, "Failed to declare a queue")
	q2, err  := ch.QueueDeclare(
		"max_nums",			// 队列名称
		false,			// 是否持久化
		false,		// 是否自动删除
		false,			// 是否独立
		false,nil,
	)
	FailOnError(err, "Failed to declare a queue")
	// 发送消息到队列中
	err = ch.Publish(
		"",     // exchange
		q0.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(a),
		})
	FailOnError(err, "Failed to publish a message")
	err = ch.Publish(
		"",     // exchange
		q1.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(b),
		})
	FailOnError(err, "Failed to publish a message")
	err = ch.Publish(
		"",     // exchange
		q2.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(c),
		})
	FailOnError(err, "Failed to publish a message")
	fmt.Println("send message success\n")

}
// 帮助函数检测每一个amqp调用
func FailOnError(err error, msg string)  {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

