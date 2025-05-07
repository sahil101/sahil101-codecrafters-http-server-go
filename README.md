# HTTP Server in Go

This repository contains an implementation of an HTTP/1.1 server in Go. The server is capable of handling multiple client connections concurrently and supports basic HTTP request parsing and response handling.

## Features
- Handles HTTP/1.1 requests
- Supports concurrent connections using goroutines
- Parses HTTP request headers, including `User-Agent`
- Responds with appropriate HTTP status codes and headers

## Getting Started

### Prerequisites
- Go (version 1.24 or later) must be installed on your system.

### Running the Server
To run the HTTP server locally:

```sh
./your_program.sh
```

This will start the server on `localhost:4221`. You can test it using tools like `curl` or a web browser.

### Example Request
```sh
curl -v http://localhost:4221/user-agent -H "User-Agent: apple/pear"
```

## Project Structure
- `app/main.go`: Entry point for the HTTP server.
- `app/parser/request_parser.go`: Handles HTTP request parsing.
- `app/response/response.go`: Handles HTTP response construction and sending.
- `your_program.sh`: Script to run the server.

## Design Patterns Used

This project incorporates several design patterns and best practices to ensure clean, maintainable, and efficient code:

### 1. **Factory Pattern**
- The `NewHTTPResponse` function in `response/response.go` acts as a factory for creating HTTP response objects. This encapsulates the creation logic and ensures consistency in response construction.

### 2. **Decorator Pattern (Conceptual)**
- While Go does not have native support for decorators, higher-order functions are used to add functionality, such as logging or error handling, to existing functions. This pattern can be extended to wrap request handlers with additional behavior.

### 3. **Concurrency Pattern**
- The server uses goroutines to handle multiple client connections concurrently. This lightweight threading model allows the server to scale efficiently and handle many connections simultaneously.

### 4. **Single Responsibility Principle**
- The project is structured to ensure that each component has a single responsibility:
  - `main.go`: Manages the server lifecycle and connection handling.
  - `parser/request_parser.go`: Handles HTTP request parsing.
  - `response/response.go`: Manages HTTP response creation and sending.

### 5. **Error Handling**
- Errors are handled gracefully throughout the code, with appropriate HTTP status codes returned to the client in case of invalid requests or server errors.

### 6. **Modular Design**
- The project is divided into distinct modules (`parser`, `response`, etc.), making it easier to maintain and extend.

These patterns and principles ensure that the project is robust, scalable, and easy to understand.

## Additional Notes
- The server uses goroutines to handle multiple client connections efficiently.
- Refer to the [HTTP/1.1 specification](https://www.w3.org/Protocols/rfc2616/rfc2616.html) for details on request and response formats.

## Resources
- [Go Documentation](https://golang.org/doc/)
- [HTTP/1.1 Specification](https://www.w3.org/Protocols/rfc2616/rfc2616.html)

Happy coding!
