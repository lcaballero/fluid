package examples

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lcaballero/fluid/calls"
)

func Ex() {
	calls.SearchAll().Do().Out()
}

func All() {
	calls.Pretty().Do().Out()
	calls.Count().Do().Out()
	calls.MatchAll().Do().Out()
	calls.PutEmployee().Do().Out()
	calls.FindEmployee().Do().Out()
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
