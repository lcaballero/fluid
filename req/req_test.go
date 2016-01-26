package req

import (
	"fmt"
	"testing"

	"net/http"
	"net/url"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReq(t *testing.T) {

	Convey("Default URL should be Loopback", t, func() {
		r := NewReq(&http.Request{
			URL: &url.URL{
				Host: "127.0.0.1:80",
			},
		})
		So(r.req.URL.Host, ShouldEqual, fmt.Sprintf("%s:%d", Loopback, HttpPort))
	})

	Convey("Default GET method", t, func() {
		r := NewReq(&http.Request{
			Method: "GET",
		})
		So(r.req.Method, ShouldEqual, GET)
	})

	Convey("Default req.Req should not be nil", t, func() {
		r := NewReq(&http.Request{})
		So(r.req, ShouldNotBeNil)
	})
}
