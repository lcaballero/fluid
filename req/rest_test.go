package req

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRest(t *testing.T) {

	Convey("rest should have host:port/_count?query", t, func() {
		s := NewRest().Get().OnLoopback(9200).Join("_count").Flag("pretty").String()
		So(s, ShouldEqual, "GET http://127.0.0.1:9200/_count?pretty")
	})

	Convey("rest should have host:port/?query", t, func() {
		s := NewRest().Get().OnLoopback(9200).Flag("pretty").String()
		So(s, ShouldEqual, "GET http://127.0.0.1:9200?pretty")
	})

	Convey("rest request should have method GET", t, func() {
		s := NewRest().Get().OnLoopback(9200)
		So(s.HttpMethod, ShouldEqual, GET)
	})

	Convey("rest request should have loopback host", t, func() {
		s := NewRest().Get()
		So(s.Url.Host, ShouldEqual, "127.0.0.1:80")
	})

	Convey("rest request should have loopback:port host", t, func() {
		s := NewRest().Get().OnLoopback(9200)
		So(s.Url.Host, ShouldEqual, "127.0.0.1:9200")
	})

	Convey("Un-Ended Rest should not have a Req", t, func() {
		s := NewRest().Get().OnLoopback(9200)
		So(s.Req, ShouldBeNil)
	})

}
