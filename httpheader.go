package main

import (
	"fmt"
)

// General Headers
// Client Request Headers
// Server Response Headers
// Entity Headers
type HTTPHeader struct {
	Key   string
	Value string
}

var (
	ServerHeader        HTTPHeader = HTTPHeader{Key: "Server"}
	DateHeader          HTTPHeader = HTTPHeader{Key: "Date"}
	ContentLengthHeader HTTPHeader = HTTPHeader{Key: "Content-Length"}
	ContentTypeHeader   HTTPHeader = HTTPHeader{Key: "Content-Type"}
	ConnectionHeader    HTTPHeader = HTTPHeader{Key: "Connection"}
)

func (h *HTTPHeader) String() string {
	return fmt.Sprintf("%s: %s", h.Key, h.Value)
}
