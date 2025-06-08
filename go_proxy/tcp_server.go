package main

import (
	"fmt"
	"net"
)

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

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the incoming data
	fmt.Printf("Received: %s", buf)

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
