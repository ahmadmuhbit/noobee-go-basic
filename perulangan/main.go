package main

import "fmt"

func main() {
	// Contoh for loop increment
	for counter := 0; counter < 5; counter++ {
		fmt.Println("Counter", counter)
	}

	fmt.Println("Counter berhenti")

	// Contoh for loop decrement
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	// Contoh nested loop
	a := [2]string{"buah", "jus"}
	b := [3]string{"jambu", "mangga", "alpukat"}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			fmt.Println(a[i], b[j])
		}
	}

	// Contoh for range
	// Slice
	fruits := []string{"Grape", "Banana", "Apple"}
	for index, fruit := range fruits {
		fmt.Println(index, fruit)
	}

	// Map
	persons := map[string]string{
		"name": "NooBee",
		"job":  "Programmer",
	}

	for _, person := range persons {
		fmt.Println(person)
	}

	// Infinite loop
	// for {
	// 	fmt.Println("ini ga bakal berhenti!")
	// }

	// Contoh break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("Selesai")

	// Contoh continue
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}

	fmt.Println("Selesai")

}
