package main

import "fmt"

func main() {
	// Mendapatkan panjang slice
	animals := []string{"Cat", "Dog", "Cow"}
	fmt.Println("Panjang slice animals:", len(animals))

	fmt.Println()

	// Mendapatkan kapasitas slice
	anim1 := animals[0:2]
	anim2 := animals[1:3]

	fmt.Println("====[ANIMALS]====")
	fmt.Println(animals)
	fmt.Println("Capacity:", cap(animals))
	fmt.Println("Len:", len(animals))
	fmt.Println()

	fmt.Println("====[ANIM1]====")
	fmt.Println(anim1)
	fmt.Println("Capacity:", cap(anim1))
	fmt.Println("Len:", len(anim1))
	fmt.Println()

	fmt.Println("====[ANIM2]====")
	fmt.Println(anim2)
	fmt.Println("Capacity:", cap(anim2))
	fmt.Println("Len:", len(anim2))
	fmt.Println()

	anim3 := animals[0:1]
	anim4 := animals[2:3]

	fmt.Println("====[ANIM3]====")
	fmt.Println(anim3)
	fmt.Println("Capacity:", cap(anim3))
	fmt.Println("Len:", len(anim3))
	fmt.Println()

	fmt.Println("====[ANIM4]====")
	fmt.Println(anim4)
	fmt.Println("Capacity:", cap(anim4))
	fmt.Println("Len:", len(anim4))
	fmt.Println()

	// Menambah elemen ke dalam slice
	fruits := []string{"Apple", "Banana", "Orange"}
	fruitSmall := fruits[0:2] // Apple, Banana

	fruit1 := append(fruits, "Grape")
	fruit2 := append(fruitSmall, "Avocado")

	fmt.Println("====FRUITS====")
	fmt.Println("Cap:", cap(fruits))
	fmt.Println("Len:", len(fruits))
	fmt.Println("Data:", fruits)
	fmt.Println()

	fmt.Println("====FRUIT SMALL====")
	fmt.Println("Cap:", cap(fruitSmall))
	fmt.Println("Len:", len(fruitSmall))
	fmt.Println("Data:", fruitSmall)
	fmt.Println()

	fmt.Println("====FRUIT 1====")
	fmt.Println("Cap:", cap(fruit1))
	fmt.Println("Len:", len(fruit1))
	fmt.Println("Data:", fruit1)
	fmt.Println()

	fmt.Println("====FRUIT 2====")
	fmt.Println("Cap:", cap(fruit2))
	fmt.Println("Len:", len(fruit2))
	fmt.Println("Data:", fruit2)
	fmt.Println()
}
