package calls

import (
	"fluid/req"
	"fmt"
	"os"
)

const DEFAULT_REST = "http://127.0.0.1:9200"

type Calls struct {
	Name string
}

func Rest(name string) *Calls {
	return &Calls{
		Name: name,
	}
}

func Pretty() *req.Req {
	return req.NewRest("pretty").OnLoopback(9200).Add("pretty", "").Get().ToReq()
}

func Shutdown() *req.Req {
	return req.NewRest("shutdown").OnLoopback(9200).Path("_shutdown").Post().ToReq()
}

func Count() *req.Req {
	return req.NewRest("count").OnLoopback(9200).
	Path("_count").Add("pretty", "").Get().Data(`
{
	"query": {
		"match_all": {}
	}
}
`).ToReq()
}

func MatchAll() *req.Req {
	return req.NewRest("match all").OnLoopback(9200).
	Path("_count").Add("pretty", "").Get().Data(`
{
	"query": {
		"match_all": {}
	}
}
`).ToReq()
}

func FindEmployee() *req.Req {
	q := `GET /megacorp/employee/%s`
	id := "1"

	if len(os.Args) >= 2 && os.Args[1] != "" {
		id = os.Args[1]
	}

	code := fmt.Sprintf(q, id)
	r, err := req.Parse(code)
	if err != nil {
		fmt.Println(err)
	}
	return r.OnLoopback(9200).ToReq()
}

func PutEmployee() *req.Req {
	s := `PUT /megacorp/employee/%s
{
	"first_name": "John",
	"last_name": "Smith",
	"age": 25,
	"about": "I love to go rock climbing",
	"interests": [ "sports", "music" ]
}
`
	id := "1"
	if len(os.Args) >= 2 && os.Args[1] != "" {
		id = os.Args[1]
	}
	code := fmt.Sprintf(s, id)
	r,err := req.Parse(code)
	if err != nil {
		fmt.Println(err)
	}
	return r.OnLoopback(9200).ToReq()
}

