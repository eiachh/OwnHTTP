package main

import (
	"errors"
	"strings"
)

type HTTPParam struct {
	Key   string
	Value string
}

type HTTPMethod int
type HTTPVersion string

const (
	GET HTTPMethod = iota //0
	POST
	PUT
	DELETE
	HEAD
	CONNECT
	OPTIONS
	TRACE
)

var httpMethodNames = [...]string{
	"GET",
	"POST",
	"PUT",
	"DELETE",
	"HEAD",
	"CONNECT",
	"OPTIONS",
	"TRACE",
}

const (
	HTTP1_0 HTTPVersion = "HTTP/1.0"
	HTTP1_1 HTTPVersion = "HTTP/1.1"
	HTTP2   HTTPVersion = "HTTP/2"
	HTTP3   HTTPVersion = "HTTP/3"
)

func (m HTTPMethod) String() string {
	if m < GET || m > DELETE {
		return "UNKNOWN"
	}
	return httpMethodNames[m]
}

func StrToMethod(str string) HTTPMethod {
	for ind, method := range httpMethodNames {
		if method == str {
			return HTTPMethod(ind)
		}
	}
	return -1
}

type Request struct {
	Method      HTTPMethod
	Root        string
	HTTPVersion string

	Params []HTTPParam

	Headers []HTTPHeader
	Body    string
}

func NewRequest(str string) (*Request, error) {
	req := Request{}
	isBody := false

	lines := strings.Split(str, "\n")
	for ind, line := range lines {
		line = strings.Trim(line, "\r")

		// REQUEST LINE
		if ind == 0 {
			splitted := strings.Split(line, " ")
			if len(splitted) != 3 {
				return nil, errors.New("bad request line")
			}
			req.Method = StrToMethod(splitted[0])
			req.Root, req.Params = ParseRoot(splitted[1])
			req.HTTPVersion = splitted[2]

			continue
		}

		// SEPARATOR
		if line == "" && !isBody {
			isBody = true
			continue
		}

		// BODY or HEADER
		if isBody {
			req.Body += line + "\r\n"
		} else {
			keyVal := strings.Split(line, ": ")
			req.Headers = append(req.Headers, HTTPHeader{Key: keyVal[0], Value: keyVal[1]})
		}
	}

	req.Body = strings.Trim(req.Body, "\r\n")
	return &req, nil
}

func ParseRoot(root string) (string, []HTTPParam) {
	retHttpParams := []HTTPParam{}

	rootNParams := strings.Split(root, "?")
	if len(rootNParams) < 2 {
		return rootNParams[0], []HTTPParam{}
	}

	params := strings.Split(rootNParams[1], "&")
	for _, param := range params {
		keyVal := strings.Split(param, "=")
		retHttpParams = append(retHttpParams, HTTPParam{Key: keyVal[0], Value: keyVal[1]})
	}

	return rootNParams[0], retHttpParams
}
