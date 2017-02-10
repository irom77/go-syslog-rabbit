package main

import (
	"net"
	"fmt"
	"log"
	"github.com/IrekRomaniuk/go-syslog-rabbit/rabbit"
	//"bytes"
	//"encoding/gob"
)

var (
	url = "amqp://guest:guest@192.168.3.45:5672"
)

func main() {
	conn, ch := rabbit.GetChannel(url)
	defer conn.Close()
	defer ch.Close()

	dataQueue := rabbit.GetQueue("threat", ch)
	
	ln, _ := listenUDP("localhost", "6000" )
	defer ln.Close()

	//buf := new(bytes.Buffer)
	//enc := gob.NewEncoder(buf)
	for {
	//go func() {
		data := handleUDPConnection(ln)
		/*message := rabbit.Message{
			Value: data,
		}*/
		//buf.Reset()
		//enc.Encode(message)
		rabbit.Publish(data, ch, dataQueue)
	//}()
	}
}

func listenUDP(hostName, portNum string) (*net.UDPConn, error) {
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

func handleUDPConnection(conn *net.UDPConn) string {

	// here is where you want to do stuff like read or write to client

	buffer := make([]byte, 1024)

	n, _, err := conn.ReadFromUDP(buffer)

	//fmt.Println("UDP client : ", addr)
	fmt.Println("Received from UDP client :  ", string(buffer[:n]))

	if err != nil {
		log.Fatal(err)
	}

	// write message back to client
	//message := []byte("Hello UDP client!")
	//_, err = conn.WriteToUDP(message, addr)

	if err != nil {
		log.Println(err)
	}
	return string(buffer[:n])

}