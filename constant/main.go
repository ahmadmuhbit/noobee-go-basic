package main

import "fmt"

// Membuat constant di luar fungsi
const LENGTH int = 10

func main() {
	// Membuat constant di dalam fungsi
	const width = 5
	const PI = 3.14

	fmt.Println(LENGTH)
	fmt.Println(width)
	fmt.Println(PI)

	// Deklarasi multiple constant
	const (
		X int = 10
		Y     = 3.14
		Z     = "Hello World"
	)

	fmt.Println(X)
	fmt.Println(Y)
	fmt.Println(Z)
}
