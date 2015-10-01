package req

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
)


func TestRest(t *testing.T) {

	Convey("rest should have host:port/_count?query", t, func() {
		s := NewRest().Get().OnLoopback(9200).Path("_count")

		fmt.Println("s.Url.Path", s.Url.Path)
		fmt.Println("s.Url.RawPath", s.Url.RawPath)
		fmt.Println("s.Url.Host", s.Url.Host)
		fmt.Println("s.Url.RawQuery", s.Url.RawQuery)
		// This call parses the RawQuery every time to produce the Values
		// map for query parameters -- TODO: do less work to get/set pairs
		fmt.Println("s.Url.Query()", s.Url.Query())
	})

	Convey("rest should have host:port/_count?query", t, func() {
		s := NewRest().Get().OnLoopback(9200).Path("_count").Flag("pretty").String()
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
}

