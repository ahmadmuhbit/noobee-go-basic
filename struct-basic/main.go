package main

import "fmt"

// Membuat struct
type Fruit struct {
	Name   string
	Weight int
	// Price  int
}

func main() {
	// Penerapan struct 1
	var fruit1 Fruit
	fruit1.Name = "Apple"
	fruit1.Weight = 1

	fmt.Println(fruit1)
	fmt.Println(fruit1.Name)
	fmt.Println(fruit1.Weight)

	// Penerapan struct 2
	fruit2 := Fruit{
		Name:   "Mango",
		Weight: 2,
	}

	fmt.Println(fruit2)
	fmt.Println(fruit2.Name)
	fmt.Println(fruit2.Weight)

	// Penerapan struct 3
	var fruit3 = Fruit{
		Name:   "Banana",
		Weight: 3,
	}

	fmt.Println(fruit3)
	fmt.Println(fruit3.Name)
	fmt.Println(fruit3.Weight)

	// Penerapan struct 4
	fruit4 := Fruit{"Lemon", 4}

	fmt.Println(fruit4)
	fmt.Println(fruit4.Name)
	fmt.Println(fruit4.Weight)
}
