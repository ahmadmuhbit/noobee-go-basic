package controllers

import (
	"belajar-package/models"
	"fmt"
)

type ContactController interface {
	DisplayMyContact()
}

type Contact struct {
	Contact models.Contact
}

func NewContactController(contact *models.Contact) ContactController {
	return &Contact{
		Contact: *contact,
	}
}

func (c *Contact) DisplayMyContact() {
	fmt.Println("Called from contact controller ...")
	fmt.Println("My Phone", c.Contact.Phone)
	fmt.Println("My Instagram", c.Contact.Instagram)
}
