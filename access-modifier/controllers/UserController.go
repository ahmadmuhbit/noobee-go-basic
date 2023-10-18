package controllers

import (
	"belajar-package/models"
	"fmt"
)

type UserController interface {
	sayHello()
	SetContact(Contact *models.Contact)
	DisplayContact()
}

type User struct {
	User models.User
}

func NewUserController(user *models.User) UserController {
	return &User{User: *user}
}

func (u *User) sayHello() {
	fmt.Println("Hello,", u.User.Name)
}

func (u *User) SetContact(contact *models.Contact) {
	u.User.Contact = *contact
}

func (u *User) DisplayContact() {
	fmt.Println("Called from user controller ...")
	fmt.Println("My Phone", u.User.Contact.Phone)
	fmt.Println("My Instagram", u.User.Contact.Instagram)
}
