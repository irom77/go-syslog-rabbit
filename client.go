package main

import "github.com/IrekRomaniuk/go-syslog-rabbit/rabbit"

func main() {
	conn, ch := rabbit.GetChannel("amqp://guest:guest@192.168.3.45:5672")
	defer conn.Close()
	defer ch.Close()

	dataQueue := rabbit.GetQueue("hello", ch)
	for {
		rabbit.Subscribe(ch, dataQueue)
	}
	//
}
