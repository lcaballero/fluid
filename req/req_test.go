package req
    
import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
)

func TestName(t *testing.T) {

	Convey("Default URL should be Loopback", t, func() {
		r := NewReq()
		So(r.Req.URL.Host, ShouldEqual, fmt.Sprintf("%s:%d", Loopback, HttpPort))
	})

	Convey("Default URL should be Loopback", t, func() {
		r := NewReq()
		So(r.Req.URL.Host, ShouldEqual, fmt.Sprintf("%s:%d", Loopback, HttpPort))
	})

	Convey("Default req.body should not be nil", t, func() {
		r := NewReq()
		So(r.Body, ShouldNotBeNil)
	})

	Convey("Default GET method", t, func() {
		r := NewReq()
		So(r.Req.Method, ShouldEqual, GET)
	})

	Convey("Default req.Req should not be nil", t, func() {
		r := NewReq()
		So(r.Req, ShouldNotBeNil)
	})
}


