package main

import (
	"fmt"
	"strconv"
)

type Response struct {
	HttpVersion HTTPVersion
	HttpCode    string
	Headers     []HTTPHeader

	Body string
}

func (resp *Response) String() string {
	strBuilder := fmt.Sprintf("%s %s %s\r\n", resp.HttpVersion, resp.HttpCode, "Created")
	for _, header := range resp.Headers {
		if header.Key != ContentLengthHeader.Key {
			strBuilder += header.String() + "\r\n"
		}
	}

	if resp.Body != "" {
		currContLengthHeader := ContentLengthHeader
		currContLengthHeader.Value = strconv.Itoa(len(resp.Body))
		strBuilder += currContLengthHeader.String() + "\r\n"

		strBuilder += "\r\n" + resp.Body
	}

	return strBuilder
}
