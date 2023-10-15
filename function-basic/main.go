package main

import "fmt"

// Membuat function
func cetak() {
	fmt.Println("Fungsi cetak() dipanggil")
	hello()
}

func hello() {
	fmt.Println("Hello World!")
}

// Parameter
func getName(name string) {
	fmt.Println(name)
}

func cetakNama(text string) {
	fmt.Println(text)
}

func cetakUmur(age int) {
	fmt.Println(age)
}

// Function parameter tipe data interface
func cetakNamaUmur(arg interface{}) {
	fmt.Println(arg)
}

// Multiple parameter
func getPeople(myName string, myAge int) {
	fmt.Println("Nama saya", myName, "Umur saya", myAge, "tahun")
}

// Function return value
func getSum(a, b int) int {
	return a + b
}

func combineString(str1 string, str2 string) string {
	return str1 + " " + str2
}

// Named return value
func addNumbers(x int, y int) (result int) {
	result = x + y
	return
}

func addNumbers2(x string) {

}

func main() {
	// Memanggil function
	cetak()
	// cetak() // Bisa dipanggil berulang-ulang
	// cetak() // Bisa dipanggil berulang-ulang

	// Parameter
	getName("NooBee")

	var names = []string{"Rey", "Jo", "NooBee"}
	var ages = []int{10, 13, 20}

	// for _, name := range names {
	// 	cetakNama(name)
	// }

	// for _, age := range ages {
	// 	cetakUmur(age)
	// }

	for _, name := range names {
		cetakNamaUmur(name)
	}

	for _, age := range ages {
		cetakNamaUmur(age)
	}

	// Multiple parameter
	getPeople("Noobee Id", 26)

	// Function return value
	fmt.Println(getSum(10, 20))
	fmt.Println(combineString("Hello", "World"))

	// Named return value
	total := addNumbers(20, 30)
	fmt.Println(total)
}
