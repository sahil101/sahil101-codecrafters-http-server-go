package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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
		// Handle the connection in a separate function
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read the request
	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
		return
	}

	// Parse the request line
	requestLine = strings.TrimSpace(requestLine)
	if strings.HasPrefix(requestLine, "GET ") {
		splitRequest := strings.Split(requestLine, " ")
		path := splitRequest[1]
		if path == "/" {
			response := "HTTP/1.1 200 OK\r\n\r\n"
			conn.Write([]byte(response))
		} else if strings.HasPrefix(path, "/echo") {
			message := strings.Split(path, "/echo/")[1]
			response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(message), message)
			conn.Write([]byte(response))

		} else {
			response := "HTTP/1.1 404 Not Found\r\n\r\n"
			conn.Write([]byte(response))
		}
	}
}
