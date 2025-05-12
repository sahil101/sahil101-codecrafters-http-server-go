package parser

import (
	"errors"
	"strings"
)

type HTTPRequest struct {
	Method        string
	Path          string
	UserAgent     string
	ContentType   string
	ContentLength string
	Body          string
}

// ParseRequest parses the HTTP request and returns an HTTPRequest struct
func ParseRequest(request string) (HTTPRequest, error) {
	lines := strings.Split(request, "\r\n") // Split by CRLF to handle headers properly
	if len(lines) == 0 {
		return HTTPRequest{}, errors.New("empty request")
	}

	// Parse the request line (first line)
	requestLine := lines[0]
	if requestLine == "" {
		return HTTPRequest{}, errors.New("missing request line")
	}

	parts := strings.Split(requestLine, " ")
	if len(parts) < 3 {
		return HTTPRequest{}, errors.New("invalid request line format")
	}

	method := parts[0]
	path := parts[1]

	// Parse headers to extract User-Agent, Content-Type, and Content-Length
	userAgent := ""
	contentType := ""
	contentLength := ""
	headersEndIndex := 0
	for i, line := range lines[1:] {
		if line == "" { // End of headers
			headersEndIndex = i + 1
			break
		}
		if strings.HasPrefix(line, "User-Agent:") {
			userAgent = strings.TrimSpace(strings.TrimPrefix(line, "User-Agent:"))
		} else if strings.HasPrefix(line, "Content-Type:") {
			contentType = strings.TrimSpace(strings.TrimPrefix(line, "Content-Type:"))
		} else if strings.HasPrefix(line, "Content-Length:") {
			contentLength = strings.TrimSpace(strings.TrimPrefix(line, "Content-Length:"))
		}
	}

	// Extract the body (if any)
	body := ""
	if headersEndIndex < len(lines) {
		body = strings.Join(lines[headersEndIndex:], "\r\n")
	}

	return HTTPRequest{
		Method:        method,
		Path:          path,
		UserAgent:     userAgent,
		ContentType:   contentType,
		ContentLength: contentLength,
		Body:          body,
	}, nil
}
