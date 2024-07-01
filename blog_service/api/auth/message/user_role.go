package message

type UserRole struct {
	User  *User    `json:"user"`
	Roles []string `json:"roles"`
}

func NewUserRole(user *User, roles ...string) *UserRole {
	return &UserRole{
		User:  user,
		Roles: roles,
	}
}
