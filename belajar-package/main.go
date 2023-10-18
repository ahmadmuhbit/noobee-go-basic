package main

import (
	"belajar-package/controllers"
	"belajar-package/models"
)

func main() {
	userModel := models.NewUser("NooBeeID")
	contactModel := models.NewContact("081285556666", "@noobeeid")

	user := controllers.NewUserController(userModel)
	contact := controllers.NewContactController(contactModel)

	user.SayHello()
	user.SetContact(contactModel)

	contact.DisplayMyContact()
	user.DisplayContact()
}
