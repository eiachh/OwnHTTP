package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var (
	response Response
	body     string
)

func Init() {
	body = `<html>
   <body>
      <h1>Hello, World!</h1>
   </body>
</html>`

	formattedTime := time.Now().UTC().Format(time.RFC1123)

	headers := []HTTPHeader{
		{Key: ServerHeader.Key, Value: "CustomServer 1.0"},
		{Key: DateHeader.Key, Value: formattedTime[:len(formattedTime)-3] + "GMT"},
		{Key: ContentTypeHeader.Key, Value: "text/html"},
		{Key: ConnectionHeader.Key, Value: "Closed"},
	}

	response = Response{
		HttpVersion: HTTP1_1,
		HttpCode:    "201",
		Headers:     headers,

		Body: body,
	}
}

func main() {
	Init()

	// Define the address and port to bind to
	address := "127.0.0.1:6969"

	// Listen on the specified address and port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Failed to bind to port:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening on", address)

	for {
		// Accept incoming client connections
		fmt.Println("WAITING")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}

		// Handle client connection
		go handleConnection(conn)
	}
}

// Handle an incoming connection
func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	requestStr := string(buf[:n])
	fmt.Printf("Received: %s\n", requestStr)

	parsedReq, err := NewRequest(requestStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("My read req: ")
	fmt.Println(*parsedReq)
	fmt.Println("My response: ")
	fmt.Println(response.String())

	conn.Write([]byte(response.String()))
}
