package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/parser"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
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
		// Handle the connection in a separate goroutine
		// Start a new goroutine to handle the connection concurrently.
		// This allows the server to handle multiple client connections at the same time.
		// Each connection is processed independently in its own lightweight thread (goroutine).
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read the request
	reader := bufio.NewReader(conn)
	// Read the full request
	request := ""
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading request: ", err.Error())
			return
		}
		if line == "\r\n" { // End of headers
			break
		}
		request += line
	}
	// Parse the full request
	httpRequest, err := parser.ParseRequest(strings.TrimSpace(request))
	if err != nil {
		resp := response.NewHTTPResponse(400, "Bad Request", response.Headers{ContentType: "text/plain"}, "Bad Request")
		resp.Send(conn)
		return
	}

	// Handle the request and send appropriate response
	if httpRequest.Method == "GET" {
		if httpRequest.Path == "/" {
			handleRoot(conn)
		} else if strings.HasPrefix(httpRequest.Path, "/echo/") {
			message := strings.TrimPrefix(httpRequest.Path, "/echo/")
			handleEcho(conn, message)
		} else if strings.HasPrefix(httpRequest.Path, "/user-agent") {
			handleUserAgent(conn, httpRequest.UserAgent)
		} else {
			handleNotFound(conn)
		}
	}
}

func handleUserAgent(conn net.Conn, userAgent string) {
	contentLength := fmt.Sprintf("%d", len(userAgent))
	rest := response.NewHTTPResponse(200, "OK", response.Headers{
		ContentType:   "text/plain",
		ContentLength: contentLength,
	}, userAgent)
	rest.Send(conn)
}

func handleRoot(conn net.Conn) {
	rest := response.NewHTTPResponse(200, "OK", response.Headers{}, "")
	rest.Send(conn)
}

func handleEcho(conn net.Conn, message string) {
	contentLength := fmt.Sprintf("%d", len(message))
	rest := response.NewHTTPResponse(200, "OK", response.Headers{
		ContentType:   "text/plain",
		ContentLength: contentLength,
	}, message)
	rest.Send(conn)
}

func handleNotFound(conn net.Conn) {
	rest := response.NewHTTPResponse(404, "Not Found", response.Headers{}, "")
	rest.Send(conn)
}
