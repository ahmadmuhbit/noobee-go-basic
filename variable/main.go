package main

import "fmt"

var x int = 1

// y := 2 // Tidak bisa di luar fungsi

func main() {
	// Membuat variable
	var myName string = "NooBee Id"

	var yourName string
	yourName = "Ubit"

	var nickName = "Noobee"

	myAge := 22

	fmt.Println(myName)
	fmt.Println(yourName)
	fmt.Println(nickName)
	fmt.Println(myAge)

	// Variable di luar fungsi
	fmt.Println(x)
	// fmt.Println(y)

	// Deklarasi variable tanpa nilai awal
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Deklarasi multiple variable
	var d, e, f, g = 1, 2, 3, "Hello"

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// Deklarasi multiple variable dalam blok
	var (
		firstName string = "Noobee"
		lastName         = "Id"
		height    int    = 170
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(height)

}
