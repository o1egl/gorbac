package gorbac

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPermission(t *testing.T) {
	Convey("Given permissions", t, func() {
		profile1 := NewPermission("profile")
		profile2 := NewPermission("profile")
		admin := NewPermission("admin")
		Convey("profile1 shoud match profile2", func() {
			So(profile1.Match(profile2), ShouldBeTrue)
		})
		Convey("profile1 should not match admin", func() {
			So(profile1.Match(admin), ShouldBeFalse)
		})
		Convey("profile1 id must be \"profile\"", func() {
			So(profile1.Id(), ShouldEqual, "profile")
		})
	})
}
