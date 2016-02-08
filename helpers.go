package gorbac

// IsGranted checks that role has permission.
func IsGranted(role Role, perm Permission) bool {
	return role.HasPermission(perm)
}

// AnyGranted checks that role has at least one permission.
func AnyGranted(role Role, perms ...Permission) bool {
	for _, p := range perms {
		if role.HasPermission(p) {
			return true
		}
	}
	return false
}

// AllGranted checks that role has all permission.
func AllGranted(role Role, perms ...Permission) bool {
	for _, p := range perms {
		if !role.HasPermission(p) {
			return false
		}
	}
	return true
}
