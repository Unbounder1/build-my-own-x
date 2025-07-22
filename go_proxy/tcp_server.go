package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

type Packet struct {
	Request    string
	Path       string
	Host       string
	Connection string
}

// Main tcp handler loop
func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	var cur_packet Packet

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error reading from connection: %v", err)
		}
		fmt.Println(line)
		lineArr := strings.Split(line, " ")
		switch lineArr[0] {
		case "GET":
			cur_packet.Request = "GET"
			cur_packet.Host = lineArr[1]
		case "POST":
			cur_packet.Request = "GET"
			cur_packet.Host = lineArr[1]
		case "Host:":
			cur_packet.Host = lineArr[1]
		}

	}

	// Print the incoming data
	// fmt.Printf("Received: %s", buf)

	response_line := "HTTP/1.1 200 OK\r\n"

	headers := "Server: Crude Server\r\nContent-Type: text/html\r\n"

	blank_line := "\r\n"

	response_body := `<html>
<body>
<h1>Request received!</h1>
<body>
</html>`

	conn.Write([]byte(response_line + headers + blank_line + response_body))
}
