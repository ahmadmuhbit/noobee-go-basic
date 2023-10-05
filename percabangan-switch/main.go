package main

import "fmt"

func main() {
	// Contoh single switch case statament
	time := "am"

	switch time {
	case "pm":
		fmt.Println("Malam")
	case "am":
		fmt.Println("Pagi")
	default:
		fmt.Println("Maaf, waktunya salah!")
	}

	// Multi case switch statement
	hari := "jumat"

	switch hari {
	case "senin", "selasa", "rabu", "kamis", "jumat":
		fmt.Println("Weekday")
	case "sabtu", "minggu":
		fmt.Println("Weekend")
	default:
		fmt.Println("Hari tidak valid")
	}
}
