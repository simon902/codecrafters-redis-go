package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	HOST = "0.0.0.0"
	PORT = "6379"
	TYPE = "tcp"
)

func requestToString(req []byte) string {
	str := string(req)

	str = strings.ReplaceAll(str, "\n", "\\n")
	return strings.ReplaceAll(str, "\r", "\\r")
}

func main() {
	fmt.Println("Logs from your program will appear here!")

	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Println("Failed to bind to port " + PORT)
		os.Exit(1)
	}

	defer listen.Close()

	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	data := make([]byte, 64)

	_, err = conn.Read(data)
	if err != nil {
		fmt.Println("Error while reading request: ", err.Error())
	}

	fmt.Printf("Received command: %s\n", requestToString(data))

	conn.Write([]byte("+PONG\r\n"))
}
