package parser

import (
	"errors"
	"strings"
)

type HTTPRequest struct {
	Method string
	Path   string
}

// ParseRequestLine parses the HTTP request line and returns an HTTPRequest struct
func ParseRequestLine(requestLine string) (HTTPRequest, error) {
	parts := strings.Split(requestLine, " ")
	if len(parts) < 2 {
		return HTTPRequest{}, errors.New("invalid request line")
	}

	method := parts[0]
	path := parts[1]

	return HTTPRequest{
		Method: method,
		Path:   path,
	}, nil
}
