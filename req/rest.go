package req

import (
	"net/url"
	"fmt"
	"path/filepath"
	"strings"
	"io"
	"os"
	"bytes"
)

//go:generate genfront --debug --input req_rest_methods.fm --output req_rest_methods.go

// A Rest structure holds data points used during any http request.  While
// the interface for a Rest instance build up the URL, Method, Query, etc
// for eventual resolution.
type Rest struct {
	url.Values
	Name string
	Url *url.URL
	HttpMethod string
	Payload string
}

const Https = "https"
const Http = "http"
const Loopback = "127.0.0.1"
const HttpPort = 80
const DefaultUrl = "http://127.0.0.1:80"

func NewRest() *Rest {
	return &Rest{
		Values: make(url.Values),
		HttpMethod: GET,
		Url: &url.URL{
			Scheme: Http,
			Host: DefaultUrl,
		},
	}
}

// Gives this Rest call a user-defined name.
func (r *Rest) Name(name string) *Rest {
	r.Name = name
	return r
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

func (r *Rest) Show(w io.Writer) *Rest {
	if w == nil {
		w = os.Stdout
	}
	fmt.Fprintln(os.Stdout, r)
	return r
}

func (r *Rest) Data(s string) *Rest {
	r.Payload = s
	return r
}

func (r *Rest) String() string {
	buf := bytes.NewBuffer([]byte{})
	fmt.Fprint(buf, r.HttpMethod, " ", r.url())

	if r.Payload != "" {
		fmt.Fprintf(buf, "\n%s", r.Payload)
	}
	return strings.Trim(buf.String(), " \t\r\n")
}

func (r *Rest) url() string {
	q := strings.Replace(r.Url.String(), "=&", "&", -1)
	if strings.HasSuffix(q, "=") {
		q = q[0:len(q) - 1]
	}
	return q
}

func (r *Rest) ToReq() *Req {
	r.Url.RawQuery = r.url()
	return NewReq().Method(r.HttpMethod).Url(r.Url).Data(r.Payload)
}

func (r *Rest) Path(p ...string) *Rest {
	root := []string{r.Url.Path}
	root = append(root, p...)
	r.Url.Path = filepath.Join(root...)
	return r
}

func (r *Rest) Flag(k string) *Rest {
	return r.Add(k, "")
}

func (r *Rest) Add(k,v string) *Rest {
	r.Values.Add(k,v)
	r.Url.RawQuery = r.Values.Encode()
	return r
}
