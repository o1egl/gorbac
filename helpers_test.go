package gorbac

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHelper(t *testing.T) {
	Convey("Given role with read, write adn edit permissions", t, func() {
		user := NewRole("user")
		readPerm := NewPermission("read")
		writePerm := NewPermission("write")
		editPerm := NewPermission("edit")
		deletePerm := NewPermission("delete")
		user.AddPermissions(readPerm, writePerm)
		Convey("Should has read and write permissions", func() {
			So(IsGranted(user, readPerm), ShouldBeTrue)
			So(IsGranted(user, writePerm), ShouldBeTrue)
			So(AllGranted(user, readPerm, writePerm), ShouldBeTrue)
		})
		Convey("Shouldn't has delete", func() {
			So(IsGranted(user, deletePerm), ShouldBeFalse)
			So(AllGranted(user, deletePerm), ShouldBeFalse)
		})
		Convey("Should has any of read and delete", func() {
			So(AnyGranted(user, readPerm, deletePerm), ShouldBeTrue)
		})
		Convey("Shouldn't has any of edit and delete", func() {
			So(AnyGranted(user, editPerm, deletePerm), ShouldBeFalse)
		})
	})
}
