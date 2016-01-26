package req

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"net/http"
	"io/ioutil"
)

//go:generate genfront front --input req_rest_methods.fm --output req_rest_methods.gen.go

// A Rest instance is a builder for an http request.
type Rest struct {
	url.Values
	Url         *url.URL
	HttpMethod  string
	Payload     []byte
	ContentType string
	ContentLenth int64
}

const (
	Https      = "https"
	Http       = "http"
	Loopback   = "127.0.0.1"
	HttpPort   = 80
	DefaultUrl = "127.0.0.1:80"
)

// NewRest provides this starting point:
//   GET http://127.0.0.1:80
//   Content-Type: text/plain
func NewRest() *Rest {
	return &Rest{
		Values:      make(url.Values),
		ContentType: "text/plain",
		HttpMethod:  GET,
		Url: &url.URL{
			Scheme: Http,
			Host:   DefaultUrl,
		},
	}
}

// Sets the Rest call on 127.0.0.1 with the given port.
func (r *Rest) OnLoopback(port int) *Rest {
	return r.Host(Loopback, port)
}

// Sets the host for this Rest call to host:port
func (r *Rest) Host(host string, port int) *Rest {
	r.Url.Host = fmt.Sprintf("%s:%d", host, port)
	return r
}

// Sets the Method for this Rest call.
func (r *Rest) Method(m string) *Rest {
	r.HttpMethod = m
	return r
}

// Write prints to the writer the results String.
func (r *Rest) Write(w io.Writer) *Rest {
	if w == nil {
		w = os.Stdout
	}
	fmt.Fprintln(w, r)
	return r
}

// Data saves the string for eventual use as a payload to the request;
// typically as POST data, be it JSON or XML or whatever.
func (r *Rest) Data(s string) *Rest {
	r.Payload = []byte(s)
	r.ContentLenth = int64(len(s))
	return r
}

// Bytes saves the given bytes for eventual use as payload for the
// request; typically as POST data.
func (r *Rest) Bytes(s []byte) *Rest {
	r.Payload = s
	r.ContentLenth = int64(len(s))
	return r
}

// IsText returns true if the content-type is set to a type of textual
// representation instead of a binary format for which it returns false.
func (r *Rest) IsText() bool {
	//TODO: check that content type is some form of 'textual' mime-type
	return r.Payload != nil && len(r.Payload) > 0 && r.ContentType != ""
}

// String provides a human readable version of the current state of
// the Rest request builder.
func (r *Rest) String() string {
	buf := bytes.NewBuffer([]byte{})
	fmt.Fprint(buf, r.HttpMethod, " ", r.url())

	if r.IsText() {
		fmt.Fprintf(buf, "\n%s", string(r.Payload))
	}
	return strings.Trim(buf.String(), " \t\r\n")
}

// URL provides a printable version of just the URL portions of this
// Rest builder.
func (r *Rest) url() string {
	r.Url.RawQuery = r.Values.Encode()
	q := strings.Replace(r.Url.String(), "=&", "&", -1)
	if strings.HasSuffix(q, "=") {
		q = q[0 : len(q)-1]
	}
	return q
}

// ToReq builds out the final Req instance.
func (r *Rest) End() (*http.Response, *bytes.Buffer, *Errors) {
	r.Url.RawQuery = r.url()
	return NewReq(r.NewRequest()).End()
}

// Out calls End and then outputs the errors if there are any, or if
// there are no errors, then the resulting body.
func (r *Rest) Out() {
	_, body, errs := r.End()
	if errs != nil {
		fmt.Println(errs)
	} else {
		fmt.Println(body)
	}
}

// Joins the values provided and uses the result as the Path for the
// resulting request URL.
func (r *Rest) Join(p ...string) *Rest {
	root := []string{r.Url.Path}
	root = append(root, p...)
	r.Url.Path = filepath.Join(root...)
	return r
}

// Flag, like a Add(k,v), add a query parameter, but without a value,
// which acts like a boolean flag: either present or absent.
func (r *Rest) Flag(k string) *Rest {
	return r.Add(k, "")
}

// Add includes the key-value pair for later addition to the request
// query parameters.
func (r *Rest) Add(k, v string) *Rest {
	r.Values.Add(k, v)
	return r
}

func (r *Rest) NewRequest() *http.Request {
	var rc io.ReadCloser = nil
	if r.ContentLenth > 0 {
		buf := bytes.NewBuffer(r.Payload)
		rc = ioutil.NopCloser(buf)
	}
	req := &http.Request{
		Method:     r.HttpMethod,
		URL:        r.Url,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Body:       rc,
		Host:       r.Url.Host,
	}
	if r.ContentLenth > 0 {
		req.ContentLength = r.ContentLenth
	}

	return req
}
