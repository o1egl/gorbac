package gorbac

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRoles(t *testing.T) {
	readPerm := NewPermission("read")
	editPerm := NewPermission("edit")
	deletePerm := NewPermission("delete")
	Convey("Given user role", t, func() {
		user := NewRole("user")
		Convey("Id should be user", func() {
			So(user.Id(), ShouldEqual, "user")
		})
		Convey("When add read and edit permissions", func() {
			user.AddPermissions(readPerm, editPerm)
			Convey("User role should have them", func() {
				So(len(user.Permissions()), ShouldEqual, 2)
				So(user.HasPermission(readPerm), ShouldBeTrue)
				So(user.HasPermission(editPerm), ShouldBeTrue)
			})
			Convey("User shouldn't has delete permission", func() {
				So(user.HasPermission(deletePerm), ShouldBeFalse)
			})
			Convey("When revoke edit permission", func() {
				user.RevokePermissions(editPerm)
				Convey("It shoudn't has edit", func() {
					So(user.HasPermission(editPerm), ShouldBeFalse)
				})
				Convey("But shoud has read", func() {
					So(user.HasPermission(readPerm), ShouldBeTrue)
				})
			})
		})
		Convey("Given admin role", func() {
			admin := NewRole("admin")
			admin.AddPermissions(deletePerm)
			Convey("When add admin parent role", func() {
				user.AddParents(admin)
				Convey("It should has admin parent", func() {
					So(user.Parents()[0].Id(), ShouldEqual, admin.Id())
				})
				Convey("It should has delete permission", func() {
					So(user.HasPermission(deletePerm), ShouldBeTrue)
				})
			})
			Convey("When remove admin parent", func() {
				user.RemoveParents(admin)
				Convey("It should not has delete permission", func() {
					So(user.HasPermission(deletePerm), ShouldBeFalse)
				})
			})
		})
	})
}
