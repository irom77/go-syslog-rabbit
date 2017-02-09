package main

import (
	"github.com/streadway/amqp"
	"log"
	"fmt"
	//"strings"
	//"encoding/json"
)

const text = `{"@timestamp":"2017-02-09T16:25:00.966Z","@version":"1","host":"10.199.107.10","message":"<141>February 09 16:25:00 R90HE73F 1,2017/09/02 16:25:00,001
901000999,threat,file,1,2017/09/02 16:25:01,1.2.3.4,2.2.2.2,0.0.0.0,0.0.0.0,G0s9J4jAU3,me,you,App test,vsys1,src,dst,ae1.100,ae2.200,LF-elk,2017/09/02 16:25:02,33891243,1
,11111,22222,0,0,0x0,tcp,test,Test,This is test only,any,low,server-to-client,5210010,0x0,10.10.10.0-10.255.255.255,10.20.20.20-10.255.255.255,0,,,,,,,,,,,,,","type":"thr
eat"}`
const text2 = "{\"Positions\": [100, 200, 300, -1]}"

type message struct {
	timestamp string // `json:@timestamp`
	version string // `json:@version`
	host string // `json:host`
	message string //`json:message`
}

type message2 struct {
	Positions []int
}

func main2() {
	go client()

	/*fmt.Printf("%s\n",text)
	bytes := []byte(text)
	var m message
	json.Unmarshal(bytes,&m)
	fmt.Println(m)

	bytes2 := []byte(text2)
	var m2 message2
	json.Unmarshal(bytes2,&m2)
	fmt.Println(m2)*/

	var a string
	fmt.Scanln(&a)
}
type PanThreatLogs struct {
	Domain,ReceiveTime,SerialNum,Type,Subtype,ConfigVersion,GenerateTime,SourceIP,DestinationIP,
	NATSourceIP,NATDestinationIP,Rule,SourceUser,DestinationUser,Application,VirtualSystem,SourceZone,DestinationZone,
	InboundInterface,OutboundInterface,LogAction,TimeLogged,SessionID,RepeatCount,SourcePort,DestinationPort,NATSourcePort,
	NATDestinationPort,Flags,Protocol,Action,URL,ThreatContentName,Category,Severity,Direction,Seqno,ActionFlags,
	SourceLocation,DestinationLocation,Cpadding_th,ContentType,Pcap_id,Filedigest,Cloud,Url_idx,User_agent,Filetype,Xff,
	Referer,Sender,Subject,Recipient,Reportid string
}



func client() {
	conn, ch, q := getQueue()

	defer conn.Close()
	defer ch.Close()

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
		log.Printf("Received message with message: %s", msg.Body)
		//log.Print(string(msg.Body))
		/*reader := csv.NewReader(strings.NewReader(string(msg.Body)))
		records, err := reader.Read()
		if err != nil {
			log.Fatal("csv ",err)
		}
		log.Print(records)*/
	}
}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.3.45:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	q, err := ch.QueueDeclare("threat",
		true, //durable bool,
		false, //autoDelete bool,
		false, //exclusivebool,
		false, //noWait bool,
		nil)   //args amqp.Table)
	failOnError(err, "Failed to declare a queue")

	return conn, ch, &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}