package main

import "fmt"

// Parameter pointer
func changeName(name *string, newName string) {
	fmt.Println("Change name from", *name, "to", newName)
	*name = newName
}

// Pointer pada method
type Car struct {
	Name  string
	Color string
}

func (c *Car) setName(newName string) {
	fmt.Println("Try to change from", c.Name, "name to =>", newName)
	c.Name = newName
}

func (c *Car) changeName2(newName string) {
	fmt.Println("Try to change from", c.Name, "name to =>", newName)
	c.Name = newName
	// fmt.Println("di dalam fungsi changeName2", c.Name)
}

func main() {
	nama := "NooBee"
	namaPointer := &nama

	fmt.Println(nama)
	fmt.Println(namaPointer)  // Mengambil alamat memor
	fmt.Println(*namaPointer) // Mengambil value

	var ptr *int
	fmt.Println(ptr)

	// Contoh penerapan pointer
	var num1 int = 5
	var num2 *int = &num1

	fmt.Println("===== Nilai Awal =====")
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(*num2)

	fmt.Println("===== Nilai Setelah num1 Diubah =====")
	num1 = 6
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(*num2)

	fmt.Println("===== Nilai Setelah num2 Diubah =====")
	*num2 = 10
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(*num2)

	// Parameter pointer
	myName := "NooBee"
	fmt.Println("Nama awal:", myName)

	changeName(&myName, "NooBeeID")
	fmt.Println("Nama setelah diubah:", myName)

	// Pointer pada method
	car := Car{Name: "BMW", Color: "White"}
	fmt.Println(car)

	car.setName("Civic")
	fmt.Println(car)

	car.changeName2("Chevrolet")
	fmt.Println(car)
}
