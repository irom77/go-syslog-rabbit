// Copyright 2017 Irek Romaniuk. All rights reserved.
/*
	Syslog server
 */
package syslogd

import (
	"log"
	"net"
	"fmt"
)

func ListenUDP(syslog string) (*net.UDPConn, error) {
	udpAddr, err := net.ResolveUDPAddr("udp4", syslog)
	if err != nil {
		return nil, err
	}

	// setup listener for incoming UDP connection
	ln, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		return nil, err
	}

	fmt.Println("UDP server " + syslog + " is up")
	return ln, nil

}

func Start(conn *net.UDPConn, debug bool) []byte {

	buffer := make([]byte, 2048)
	n, _, err := conn.ReadFromUDP(buffer)
	//fmt.Println("UDP client : ", addr)
	if debug {
		fmt.Println("Received from UDP client :  ", string(buffer[:n]))
	}
	if err != nil {
		log.Fatal(err)
	}
	return (buffer[:n])
}
