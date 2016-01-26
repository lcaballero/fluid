package examples

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lcaballero/fluid/calls"
	"github.com/lcaballero/fluid/req"
)

func Simple() {
	_, body, errs := req.NewRest().
		Host("www.google.com", 80).
		End()

	if errs != nil {
		fmt.Println(errs)
	}

	fmt.Println(body)
}

func Ex() {
	calls.SearchAll().Out()
}

func All() {
	calls.Pretty().Out()
	calls.Count().Out()
	calls.MatchAll().Out()
	calls.PutEmployee().Out()
	calls.FindEmployee().Out()
}

func BatchEmployee() {
	buf := `
{"index":{}}
{"first_name": "John","last_name": "Smith","age": 25,"about": "I love to go rock climbing","interests": [ "sports", "music" ]}


`
	cl := ioutil.NopCloser(bytes.NewBufferString(buf))

	res, err := http.Post("http://127.0.0.1:9200/bz1/business/_bulk", "json", cl)
	if err != nil {
		fmt.Println(err)
	}

	out := bytes.NewBufferString("")
	err = res.Write(out)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out.String())
}
