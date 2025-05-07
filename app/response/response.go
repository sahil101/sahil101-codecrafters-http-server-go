package response

import (
	"fmt"
	"net"
)

type HTTPResponse struct {
	StatusCode int
	StatusText string
	Headers    map[string]string
	Body       string
}

type Headers struct {
	ContentType   string
	ContentLength string
}

func NewHTTPResponse(statusCode int, statusText string, headers Headers, body string) HTTPResponse {
	if statusText == "" {
		statusText = getStatusText(statusCode)
	}
	return HTTPResponse{
		StatusCode: statusCode,
		StatusText: statusText,
		Headers: map[string]string{
			"Content-Type":   headers.ContentType,
			"Content-Length": headers.ContentLength,
		},
		Body: body,
	}
}

// Helper function to map status codes to default status texts
func getStatusText(statusCode int) string {
	switch statusCode {
	case 200:
		return "OK!"
	case 400:
		return "Bad Request"
	case 404:
		return "Not Found"
	default:
		return ""
	}
}

func (r HTTPResponse) ToString() string {
	headers := ""
	for key, value := range r.Headers {
		headers += key + ": " + value + "\r\n"
	}

	return "HTTP/1.1 " + string(r.StatusCode) + "\r\n" +
		headers + "\r\n" +
		r.Body
}

// New function to send the HTTP response over a connection
func (r *HTTPResponse) Send(conn net.Conn) {
	// Construct the status line
	statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.StatusCode, r.StatusText)

	// Construct the headers
	headers := ""
	for key, value := range r.Headers {
		headers += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	// Combine all sections: status line, headers, and response body
	response := statusLine + headers + "\r\n" + r.Body

	// Write the response to the connection
	conn.Write([]byte(response))
}
