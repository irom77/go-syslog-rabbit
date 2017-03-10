// Copyright 2017 Irek Romaniuk. All rights reserved.
/*
	Syslog to Rabbit server
 */
package main

import (
	"github.com/irom77/go-syslog-rabbit/rabbit"
	"github.com/irom77/go-syslog-rabbit/syslogd"
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"os"
)
var (
	RABVR = flag.String("r", "guest:guest@192.168.3.51:5672", "Rabbit server")
	SYSVR = flag.String("s", "0.0.0.0:12514", "Syslog server")
	DEBUG = flag.Bool("p", false, "Print debug")
	QUEUE = flag.String("q", "threat", "Name of the queue")
	version = flag.Bool("v", false, "Prints current version")
)
var (
	Version   = "No Version Provided"
	BuildTime = ""
)
func init() {
	flag.Usage = func() {
		fmt.Printf("Copyright 2017 @IrekRomaniuk. All rights reserved.\n")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *version {
		fmt.Printf("App Version: %s\nBuild Time : %s\n", Version, BuildTime)
		os.Exit(0)
	}
}

var (
	url = "amqp://"
)

func main() {
	url += *RABVR
	fmt.Println("Connecting to " + url)
	conn, ch := rabbit.GetChannel(url)
	defer conn.Close()
	defer ch.Close()

	dataQueue := rabbit.GetQueue(*QUEUE, ch)
	
	ln, _ := syslogd.ListenUDP(*SYSVR)
	// 10000 messages with freq 100 -> 10,000 rcvd - syslog on win7, rabbit on Debian
	// 10000 messages with freq 500 -> max 9,816 rcvd
	// 10000 messages with freq 1000 -> max 9,377 rcvd
	defer ln.Close()
	fmt.Printf("DEBUG %v\n%s -> %s\n", *DEBUG,*SYSVR, url)
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	for {
		//go func() {
		data := syslogd.Start(ln, *DEBUG)
		message := rabbit.Message{
			Value: data,
		}
		buf.Reset()
		enc.Encode(message)
		rabbit.Publish(buf, ch, dataQueue)
		//}()
	}
}

