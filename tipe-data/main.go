package main

import "fmt"

func main() {
	// Contoh tipe data string
	var txt string = "Ini adalah sebuah teks"
	// fmt.Println(txt)
	fmt.Printf("Value: %v, tipe data: %T\n", txt, txt)

	var txt2 string = `
	Ini adalah "teks".
	Dan ini dari 'baris baru'.
	Oke..
	`
	fmt.Println(txt2)

	// Contoh tipe data integer
	var a int16 = 30000
	var age uint8 = 22
	var bigNumber int64 = -9223372036854775807

	fmt.Println(a)
	fmt.Printf("%d\n", age)
	fmt.Printf("%d\n", bigNumber)

	// Contoh tipe data float
	var x float64 = 3.14
	fmt.Printf("%f\n", x)
	fmt.Printf("%.1f\n", x)
	fmt.Printf("%.2f\n", x)
	fmt.Println(x)

	// Contoh tipe data bool
	var isComplete bool
	var isMarried bool = true

	fmt.Println(isComplete)
	fmt.Println(isMarried)
}
