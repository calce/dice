package dice

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestId(t *testing.T) {

	Convey("Generated MAC IDs should be unique", t, func() {
		list := make(map[string]bool)
		for i:=0; i<100; i++ {
			id := LessSpin()
			So(list[id], ShouldBeFalse)
			list[id] = true
		}
	})

	Convey("Generated random IDs should be unique", t, func() {
		list := make(map[string]bool)
		id := ""
		for i:=0; i<100; i++ {
			id = Spin()
			So(list[id], ShouldBeFalse)
			list[id] = true
		}
	})

}