package main

import "fmt"

func main() {
	// Membuat array dengan keyword var dan tanda :=
	// Inisialisasi panjang data
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"
	// fruits[3] = "grape" // error, karena melebihi kapasitas array

	fmt.Println(fruits)

	a1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a1)

	// Tanpa inisialisasi panjang data
	cars := [...]string{
		"BMW",
		"Renault",
		"Audi",
		"Honda",
	}
	fmt.Println(cars)

	// Mebuat array dengan keyword make
	numbers := make([]int, 3)
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	fmt.Println(numbers)

	// Mengakses elemen array
	fmt.Println(fruits[0])
	fmt.Println(fruits[1])
	fmt.Println(fruits[2])

	// Mengubah elemen pada array
	fruits[1] = "Pisang"
	fmt.Println(fruits[1])

	// Mendapatkan panjang array
	fmt.Println(len(fruits))
	fmt.Println(len(a1))
	fmt.Println(len(cars))
	fmt.Println(len(numbers))

	// Array multidimensi
	var fruits2 = [2][3]string{
		[3]string{"Apple", "Mango", "Banana"},
		[3]string{"Orange", "Grape", "Avocado"},
	}

	fmt.Println(fruits2)
	// Mengakses elemen avocado
	fmt.Println(fruits2[1][2])
	// Mendapatkan panjang array
	fmt.Println(len(fruits2))
}
