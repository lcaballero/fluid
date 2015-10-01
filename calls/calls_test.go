package calls
   
import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestName(t *testing.T) {

	Convey("ctor should have set name", t, func() {
		name := "ctor"
		r := Rest(name)
		So(r.Name, ShouldEqual, name)
	})
}


