/*
	Syslog to Rabbit client
 */
// +build
package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/irom77/go-syslog-rabbit/rabbit"
)

var (
	cSERVR = flag.String("r", "guest:guest@192.168.3.51:5672", "Rabbit server")
	cQUEUE = flag.String("q", "threat", "Name of the queue")
	Cversion = flag.Bool("v", false, "Prints current version")
)
var (
	cVersion   = "No Version Provided"
	cBuildTime = ""
)
func init() {
	flag.Usage = func() {
		fmt.Printf("Copyright 2017 @IrekRomaniuk. All rights reserved.\n")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *Cversion {
		fmt.Printf("App Version: %s\nBuild Time : %s\n", cVersion, cBuildTime)
		os.Exit(0)
	}
}
func main() {

	conn, ch := rabbit.GetChannel("amqp://" + *cSERVR)
	defer conn.Close()
	defer ch.Close()

	dataQueue := rabbit.GetQueue(*cQUEUE, ch)
	rabbit.Subscribe(ch, dataQueue)
}
