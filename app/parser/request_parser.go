package parser

import (
	"errors"
	"strings"
)

type HTTPRequest struct {
	Method    string
	Path      string
	UserAgent string
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

	// Parse headers to extract User-Agent
	userAgent := ""
	for _, line := range lines[1:] {
		if strings.HasPrefix(line, "User-Agent:") {
			userAgent = strings.TrimSpace(strings.TrimPrefix(line, "User-Agent:"))
			break
		}
	}

	return HTTPRequest{
		Method:    method,
		Path:      path,
		UserAgent: userAgent,
	}, nil
}
