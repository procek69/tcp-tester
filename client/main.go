package main

import (
	"flag"
	"log"
	"net"
	"strconv"
	"strings"
)

const (
	message       = "Ping"
	StopCharacter = "\r\n\r\n"
)

func SocketClient(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln("Unable to connect " + ip + ":" + strconv.Itoa(port))
	}
	defer conn.Close()

	conn.Write([]byte(message))
	conn.Write([]byte(StopCharacter))
	log.Printf("Send: %s", message)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])

}

func main() {

	var targetIP = flag.String("h", "127.0.0.1", "Target IP")
	var targetPort = flag.Int("p", 8080, "Target port")
	flag.Parse()
	SocketClient(*targetIP, *targetPort)
}
