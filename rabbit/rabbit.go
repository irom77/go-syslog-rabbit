package rabbit

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)

/*func init() {
	gob.Register(Message{})
}

type Message struct {
	Value     []byte
}*/

func Publish(message string, ch *amqp.Channel,q *amqp.Queue) {
	msg := amqp.Publishing{
		//ContentType: "text/plain",
		Body:        []byte(message),
	}
	ch.Publish("", q.Name, false, false, msg)
}

func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to establish connection to message broker")
	ch, err := conn.Channel()
	failOnError(err, "Failed to get channel for connection")

	return conn, ch
}

func GetQueue(qName string, ch *amqp.Channel) *amqp.Queue {

	q, err := ch.QueueDeclare(qName,
		false, //durable bool,
		false, //autoDelete bool,
		false, //exclusivebool,
		false, //noWait bool,
		nil)   //args amqp.Table)
	failOnError(err, "Failed to declare a queue")

	return &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func Subscribe(ch *amqp.Channel, q *amqp.Queue ) {
	msgs, err := ch.Consume(
		q.Name, //queue string,
		"",     //consumer string,
		true,   //autoAck bool,
		false,  //exclusive bool,
		false,  //noLocal bool,
		false,  //noWait bool,
		nil)    //args amqp.Table)

	failOnError(err, "Failed to register a consumer")

	for msg := range msgs {
		/*r := bytes.NewReader(msg.Body)
		d := gob.NewDecoder(r)
		sd := new(Message)
		d.Decode(sd)*/
		log.Printf("Received message with message: %s", msg.Body)
	}
}
