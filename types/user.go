package types

type User struct {
	Name string
}

func NewUser(name string) *User {
	u := &User{Name: name}
	return u
}
