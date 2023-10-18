package models

type Contact struct {
	Phone     string
	Instagram string
}

func NewContact(phone, instagram string) *Contact {
	return &Contact{
		Phone:     phone,
		Instagram: instagram,
	}
}
