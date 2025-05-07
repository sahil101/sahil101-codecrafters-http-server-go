package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Start listening on port 4221
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	for {
		// Accept a connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		// Handle the connection in a separate function
		// handleConnection(conn)
	}
}

// func handleConnection(conn net.Conn) {
// 	defer conn.Close()

// 	// Read the request
// 	reader := bufio.NewReader(conn)
// 	requestLine, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error reading request: ", err.Error())
// 		return
// 	}

// 	// // Parse the request line
// 	// requestLine = strings.TrimSpace(requestLine)
// 	// if strings.HasPrefix(requestLine, "GET ") {
// 	// 	// Write a simple HTTP response
// 	// 	response := "HTTP/1.1 200 OK\r\n" +
// 	// 		"Content-Type: text/plain\r\n" +
// 	// 		"Content-Length: 13\r\n" +
// 	// 		"\r\n" +
// 	// 		"Hello, world!"
// 	// 	conn.Write([]byte(response))
// 	// } else {
// 	// 	// Respond with 400 Bad Request for unsupported methods
// 	// 	response := "HTTP/1.1 400 Bad Request\r\n" +
// 	// 		"Content-Length: 0\r\n" +
// 	// 		"\r\n"
// 	// 	conn.Write([]byte(response))
// 	// }
// }
