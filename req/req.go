package req

import (
	"net/http"
	"bytes"
	"log"
	"net/url"
	"fmt"
	"io/ioutil"
)

type Req struct {
	Name string
	Description string

	// Base (for GET)
	Req *http.Request
	Res *http.Response
	err error

	// For POST
	Body *bytes.Buffer
}

func loopback() *url.URL {
	u, err := url.Parse("http://127.0.0.1:80")
	if err != nil {
		log.Fatal(err)
	}
	return u
}

// Request instance is remade because .NewRequest defaults a number of the
// Request fields in preparation to make a network http call.
func (r *Req) reMakeRequest() *Req {
	r.Req, r.err = http.NewRequest(r.Req.Method, r.Req.URL.String(), r.Req.Body)
	return r
}

func (r *Req) parseUrl(u string) *Req {
	r.Req.URL, r.err = url.Parse(u)
	return r
}

// Creates a default req with Request instance and byte buffer for body
// reader/closer.
func NewReq() *Req {
	return &Req{
		Req: &http.Request{
			Method: GET,
			URL: loopback(),
		},
		Body: bytes.NewBuffer([]byte{}),
	}
}

func (r *Req) Data(p string) *Req {
	buf := bytes.NewBufferString(p)
	cl := ioutil.NopCloser(buf)
	r.Body = buf
	r.Req.Body = cl
	return r
}

func (r *Req) Method(m string) *Req {
	r.Req.Method = m
	return r
}

func (r *Req) ParseUrl(u string) *Req {
	return r.parseUrl(u)
}

func (r *Req) Url(u *url.URL) *Req {
	r.Req.URL = u
	return r
}

func (r *Req) HasError() bool {
	return r.err != nil
}

func (r *Req) Do() *Req {
	//http.Post(url, bodyType, io.Reader)
	r.reMakeRequest()
	if r.HasError() {
		return r
	}
	r.Res, r.err = http.DefaultClient.Do(r.Req)
	if r.HasError() {
		return r
	}
	return r.ReadAll()
}

func (r *Req) Fatal() *Req {
	if r.HasError() {
		log.Fatal(r.Error())
	}
	return r
}

func (r *Req) Out() *Req {
	if r.HasError() {
		fmt.Println(r.err)
	} else {
		fmt.Println(r)
	}
	return r
}

func (r *Req) ReadAll() *Req {
	r.err = r.Res.Write(r.Body)
	if r.HasError() {
		return r
	}
	return r
}

func (r *Req) Error() error {
	return r.err
}

func (r *Req) String() string {
	return r.Body.String()
}

