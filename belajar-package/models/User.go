package models // define nama package

type User struct {
	Name    string
	Contact Contact
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}
