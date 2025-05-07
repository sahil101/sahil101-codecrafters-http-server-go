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

// ParseRequestLine parses the HTTP request line and returns an HTTPRequest struct
func ParseRequestLine(requestLine string) (HTTPRequest, error) {
	parts := strings.Split(requestLine, " ")
	if len(parts) < 2 {
		return HTTPRequest{}, errors.New("invalid request line")
	}

	method := parts[0]
	path := parts[1]

	// Extract User-Agent if present
	userAgent := ""
	for _, line := range strings.Split(requestLine, "\n") {
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
