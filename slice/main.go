package main

import "fmt"

func main() {
	// Membuat array dan slice
	iniArray1 := [4]string{"Cat", "Dog", "Chicken", "Bird"}        // Array
	iniArray2 := [...]string{"BMW", "Renault", "Audi", "Mercedes"} // Array
	iniSlice := []string{"Apple", "Banana", "Mango", "Avocado"}    // Slice

	fmt.Println(iniArray1)
	fmt.Println(iniArray2)
	fmt.Println(iniSlice)

	// Membuat slice dari array
	iniSlice2 := iniArray1[0:2]
	iniSlice3 := iniArray1[1:]
	iniSlice4 := iniArray1[:3]
	iniSlice5 := iniArray1[:]

	fmt.Println(iniSlice2)
	fmt.Println(iniSlice3)
	fmt.Println(iniSlice4)
	fmt.Println(iniSlice5)

	// Slice tipe data reference
	iniSlice2[1] = "Cow"
	fmt.Println("iniSlice2:", iniSlice2)
	fmt.Println(iniSlice3)
	fmt.Println(iniSlice4)
	fmt.Println(iniSlice5)
	fmt.Println(iniArray1)

	iniSlice6 := iniArray1[:]
	fmt.Println("Inislice6:", iniSlice6)

	// Membuat slice dengan fungsi make
	numbers := make([]int, 5, 5)
	numbers[0] = 10
	numbers[1] = 20
	fmt.Println(numbers)

	iniSlice2[1] = "Dog"
	fmt.Println("iniSlice2:", iniSlice2)
	iniSlice2[1] = "Fish"
	fmt.Println("iniSlice2:", iniSlice2)

}
