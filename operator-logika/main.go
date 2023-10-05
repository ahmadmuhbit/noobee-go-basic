package main

import "fmt"

func main() {
	// Contoh operator AND, OR, NOT
	var a = 15

	fmt.Println(a > 10 && a < 5)
	fmt.Println(a > 10 || a < 5)
	fmt.Println(!(a > 10 || a < 5))

	// Contoh penggunaan
	gender := "male"
	age := 17

	var canWork bool

	// Penggunaan AND
	if gender == "male" && age >= 18 {
		canWork = true
	} else {
		canWork = false
	}

	fmt.Println("Nilai canWork pada notasi AND:", canWork)

	// Penggunaan OR
	if gender == "male" || age >= 18 {
		canWork = true
	} else {
		canWork = false
	}

	fmt.Println("Nilai canWork pada notasi OR:", canWork)

	// Penggunaan NOT
	fmt.Println("Nilai canWork pada notasi NOT:", !(canWork))
}
