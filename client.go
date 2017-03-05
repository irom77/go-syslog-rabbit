// Copyright 2017 Irek Romaniuk. All rights reserved.
/*
	Syslog to Rabbit client
 */
package main

import (
	"github.com/irom77/go-syslog-rabbit/rabbit"
)


func main() {

	conn, ch := rabbit.GetChannel("amqp://guest:guest@192.168.3.51:5672")
	defer conn.Close()
	defer ch.Close()

	dataQueue := rabbit.GetQueue("threat", ch)
	rabbit.Subscribe(ch, dataQueue)
}
