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

func ListenUDP(hostName, portNum string) (*net.UDPConn, error) {
	service := hostName + ":" + portNum

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		return nil, err
	}

	// setup listener for incoming UDP connection
	ln, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, err
	}

	fmt.Println("UDP server up and listening on port " + portNum)
	return ln, nil

}

func Start(conn *net.UDPConn) []byte {

	buffer := make([]byte, 2048)
	n, _, err := conn.ReadFromUDP(buffer)
	//fmt.Println("UDP client : ", addr)
	//fmt.Println("Received from UDP client :  ", string(buffer[:n]))
	if err != nil {
		log.Fatal(err)
	}
	return (buffer[:n])
}
