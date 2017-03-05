package main
// Copyright 2017 Irek Romaniuk. All rights reserved.

import (
	"github.com/irom77/go-syslog-rabbit/rabbit"
	"github.com/irom77/go-syslog-rabbit/syslogd"
	"bytes"
	"encoding/gob"
)

var (
	url = "amqp://guest:guest@192.168.3.51:5672"
)

func main() {
	conn, ch := rabbit.GetChannel(url)
	defer conn.Close()
	defer ch.Close()

	dataQueue := rabbit.GetQueue("threat", ch)
	
	ln, _ := syslogd.ListenUDP("localhost", "6000")
	// 10000 messages with freq 100 -> 10,000 rcvd - syslog on win7, rabbit on Debian
	// 10000 messages with freq 500 -> max 9,816 rcvd
	// 10000 messages with freq 1000 -> max 9,377 rcvd
	defer ln.Close()

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	for {
		//go func() {
		data := syslogd.Start(ln)
		message := rabbit.Message{
			Value: data, //data["content"],
		}
		buf.Reset()
		enc.Encode(message)
		rabbit.Publish(buf, ch, dataQueue)
		//}()
	}
}

