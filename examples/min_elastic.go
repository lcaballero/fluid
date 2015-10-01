package examples
import "fluid/calls"


func Ex() {
	calls.Pretty().Do().Out()
	calls.Count().Do().Out()
	calls.MatchAll().Do().Out()
	calls.PutEmployee().Do().Out()
	calls.FindEmployee().Do().Out()
}

