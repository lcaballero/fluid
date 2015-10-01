package req

import (
	"strings"
	"bytes"
	"bufio"
	"errors"
	"io"
	"fmt"
)


type Shorthand struct {}
var httpMethods = []string{ GET, HEAD, PUT, POST, INS, DEL, }

func parseUrl(r *Rest, reader *bufio.Reader) (*Rest, error) {

	// Skips empty lines and uses the first line found as the method+req line
	lineBytes, _, _ := reader.ReadLine()
	line := string(lineBytes)
	for line == "" {
		lineBytes, _, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}
		line = string(lineBytes)
	}

	method := ""
	for _,m := range httpMethods {
		fmt.Println(m, line)
		if strings.HasPrefix(line, m) {
			method = m
			break
		}
	}
	if method == "" {
		return nil, errors.New("Did not correctly parse http method")
	}

	r.Url.RawPath = strings.Trim(line[len(method):], " \t")
	r.Url.Path = r.Url.RawPath
	return r.Method(method), nil
}

func parsePayload(r *Rest, reader *bufio.Reader) (*Rest, error) {
	line, prefix, err := reader.ReadLine()

	// Skipping any empty lines between raw path and payload
	for string(line) == "" && err == nil {
		line, _, err = reader.ReadLine()
	}

	buf := bytes.NewBufferString("")
	for err == nil && line != nil {
		buf.Write(line)
		if !prefix {
			buf.WriteString("\n")
		}
		line, prefix, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
	}

	if err != nil && err != io.EOF  {
		return nil, err
	}

	r.Payload = buf.String()

	return r, nil
}

func Parse(sh string) (*Rest, error) {
	r := NewRest("parsed")
	buf := bufio.NewReader(bytes.NewBufferString(sh))

	r, err := parseUrl(r, buf)
	if err != nil {
		return nil, err
	}

	r, err = parsePayload(r, buf)
	if err != nil {
		return nil, err
	}

	return r, nil
}