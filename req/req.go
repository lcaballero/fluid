package req

import (
	"bytes"
	"net/http"
	"fmt"
)

// A Req is an internal structure used to make the http Client
// and executing the http call.
type Req struct {
	req    *http.Request
	errors *Errors
}

// Creates a default req with Request instance and byte buffer for body
// reader/closer.
func NewReq(r *http.Request) *Req {
	return &Req{
		req:    r,
		errors: &Errors{},
	}
}

// End() carries out the request and provides the response, with the
// body read from the request, and any errors that occurred.
func (r *Req) End() (*http.Response, *bytes.Buffer, *Errors) {
	fmt.Println(r.req)
	fmt.Println(r.req.URL)
	res, err := http.DefaultClient.Do(r.req)

	r.errors.Add(err)
	if r.errors.HasError() {
		return nil, nil, r.errors
	}

	body := bytes.NewBuffer([]byte{})

	r.errors.Add(res.Write(body))
	if r.errors.HasError() {
		return nil, nil, r.errors
	}

	return res, body, nil
}
