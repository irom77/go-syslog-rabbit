package main

import (
	"fmt"
	"gopkg.in/mcuadros/go-syslog.v2"
        //"github.com/mcuadros/go-syslog"
	"log/syslog"
)

func main() {

	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.RFC5424)
	server.SetHandler(handler)
	server.ListenUDP("0.0.0.0:12514")
	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			fmt.Println(logParts)
		}
	}(channel)

	server.Wait()

}



