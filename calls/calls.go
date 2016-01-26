package calls

import (
	"fmt"
	"os"

	"github.com/lcaballero/fluid/req"
	"github.com/lcaballero/fluid/shorthand"
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

func Pretty() *req.Rest {
	return req.NewRest().OnLoopback(9200).Add("pretty", "").Get()
}

func Shutdown() *req.Rest {
	return req.NewRest().OnLoopback(9200).Join("_shutdown").Post()
}

func Count() *req.Rest {
	return req.NewRest().OnLoopback(9200).
		Join("_count").Add("pretty", "").Get().Data(`
{
	"query": {
		"match_all": {}
	}
}
`)
}

func MatchAll() *req.Rest {
	return req.NewRest().OnLoopback(9200).
		Join("_count").Add("pretty", "").Get().Data(`
{
	"query": {
		"match_all": {}
	}
}
`)
}

func FindEmployee() *req.Rest {
	q := `GET /megacorp/employee/%s`
	id := "1"

	if len(os.Args) >= 2 && os.Args[1] != "" {
		id = os.Args[1]
	}

	code := fmt.Sprintf(q, id)
	r, err := shorthand.Parse(code)
	if err != nil {
		fmt.Println(err)
	}
	return r.OnLoopback(9200)
}

func PutEmployee() *req.Rest {
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
	r, err := shorthand.Parse(code)
	if err != nil {
		fmt.Println(err)
	}
	return r.OnLoopback(9200)
}

func SearchAll() *req.Rest {
	s := `
{
  "query": {
    "match_all": {}
  }
}
`
	return req.NewRest().OnLoopback(9200).Join("_search").Add("pretty", "").Get().Data(s)
}
