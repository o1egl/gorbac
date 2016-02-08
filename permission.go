package gorbac

type NewPermissionFunc func(string) Permission

type Permission interface {
	Id() string
	Match(Permission) bool
}

type Permissions map[string]Permission

// StdPermission only checks if the Ids are fully matching.
type StdPermission struct {
	IdStr string
}

func NewPermission(id string) Permission {
	return &StdPermission{id}
}

func (p *StdPermission) Id() string {
	return p.IdStr
}

func (p *StdPermission) Match(a Permission) bool {
	return p.IdStr == a.Id()
}
