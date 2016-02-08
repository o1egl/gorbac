package gorbac

import "sync"

type Role interface {
	Id() string
	AddPermissions(...Permission) error
	HasPermission(Permission) bool
	RevokePermissions(...Permission) error
	Permissions() []Permission
	AddParents(...Role) error
	RemoveParents(...Role) error
	Parents() []Role
}

type Roles map[string]Role

// NewStdRole is the default role factory function.
// It matches the declaration to RoleFactoryFunc.
func NewRole(id string) Role {
	role := &StdRole{
		IdStr:       id,
		permissions: make(Permissions),
		parents:     make(Roles),
	}
	return role
}

// StdRole is the default role implement.
type StdRole struct {
	sync.RWMutex
	IdStr       string
	permissions Permissions
	parents     Roles
}

// Name returns the role's identity name.
func (role *StdRole) Id() string {
	return role.IdStr
}

// AddPermissions adds a permission to the role.
func (role *StdRole) AddPermissions(ps ...Permission) error {
	role.Lock()
	defer role.Unlock()
	for _, p := range ps {
		role.permissions[p.Id()] = p
	}
	return nil
}

// HasPermission returns true if the role or its parents have specific permission.
func (role *StdRole) HasPermission(p Permission) bool {
	role.RLock()
	defer role.RUnlock()
	for _, rp := range role.permissions {
		if rp.Match(p) {
			return true
		}
	}
	for _, parent := range role.parents {
		if parent.HasPermission(p) {
			return true
		}
	}
	return false
}

// RevokePermissions remove the specific permission.
func (role *StdRole) RevokePermissions(ps ...Permission) error {
	role.Lock()
	defer role.Unlock()
	for _, p := range ps {
		delete(role.permissions, p.Id())
	}
	return nil
}

// Permissions returns all permissions into a slice.
func (role *StdRole) Permissions() []Permission {
	role.RLock()
	defer role.RUnlock()
	result := make([]Permission, 0, len(role.permissions))
	for _, p := range role.permissions {
		result = append(result, p)
	}
	return result
}

// AddParents adds parent roles.
func (role *StdRole) AddParents(parents ...Role) error {
	role.Lock()
	defer role.Unlock()
	for _, parent := range parents {
		role.parents[parent.Id()] = parent
	}
	return nil
}

// RemoveParents removes parent roles.
func (role *StdRole) RemoveParents(parents ...Role) error {
	role.Lock()
	defer role.Unlock()
	for _, parent := range parents {
		delete(role.parents, parent.Id())
	}
	return nil
}

// Parents returns parent roles.
func (role *StdRole) Parents() []Role {
	role.RLock()
	defer role.RUnlock()
	var pSlice []Role
	for _, parent := range role.parents {
		pSlice = append(pSlice, parent)
	}
	return pSlice
}
