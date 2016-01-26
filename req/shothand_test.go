package req

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var rawPath = "/_count"
var template = `
%s %s
%s`

var rawJson = `{
	"query": {
		"match_all": {}
	}
}
`

func ex() string {
	return fmt.Sprintf(template, GET, rawPath, rawJson)
}

func TestShorthand(t *testing.T) {

	Convey("short hand parse should find data payload", t, func() {
		s := `PUT /megacorp/employee/%d
{
	"first_name": "John",
	"last_name": "Smith",
	"age": 25,
	"about": "I love to go rock climbing",
	"interests": [ "sports", "music" ]
}
`
		code := fmt.Sprintf(s, 1)
		restReq, err := Parse(code)
		So(err, ShouldBeNil)
		So(restReq.Url.RawPath, ShouldEqual, "/megacorp/employee/1")

		rq := restReq.ToReq()
		So(rq, ShouldNotBeNil)
		So(rq.Req.URL.RawPath, ShouldEqual, "/megacorp/employee/1")

		rq = restReq.OnLoopback(9200).ToReq()

		So(rq.Req.URL.Host, ShouldEqual, "127.0.0.1:9200")
	})

	Convey("short hand parse should find data payload", t, func() {
		rq, err := Parse(ex())
		So(err, ShouldBeNil)
		So(rq, ShouldNotBeNil)
		So(rq.Payload, ShouldEqual, rawJson)
	})

	Convey("short hand parse should find raw path", t, func() {
		rq, err := Parse(ex())
		So(err, ShouldBeNil)
		So(rq, ShouldNotBeNil)
		So(rq.Url.RawPath, ShouldEqual, rawPath)
	})

	Convey("short hand parse should find method", t, func() {
		rq, err := Parse(ex())
		So(err, ShouldBeNil)
		So(rq, ShouldNotBeNil)
		So(rq.HttpMethod, ShouldEqual, GET)
	})
}
