package main

import "fmt"

func main() {
	// Operator perbandingan
	x := 10
	y := 5

	fmt.Println("x == y:", x == y)
	fmt.Println("x != y:", x != y)
	fmt.Println("x > y:", x > y)
	fmt.Println("x < y:", x < y)
	fmt.Println("x >= y:", x >= y)
	fmt.Println("x <= y:", x <= y)

	// Percabangan if
	num := 2
	if num > 2 {
		// Kondisi ini akan dipanggil jika syarat terpenuhi
		fmt.Println("Num lebih besar dari 2")
	} else {
		// kondisi ini akan dipanggil jika
		// kondisi pertama tidak terpenuhi
		fmt.Println("Num lebih kecil dan sama dari 2")
	}

	// akan dijalankan setelah program di dalam selesai
	fmt.Println("Program selesai")

	// Percabangan if - else if - else
	nilai := 80
	var grade string

	if nilai > 80 {
		grade = "A"
	} else if nilai > 70 {
		grade = "B"
	} else if nilai > 60 {
		grade = "C"
	} else if nilai > 50 {
		grade = "D"
	} else {
		grade = "E"
	}
	fmt.Println("Nilai saya adalah:", grade)

	// Percabangan nested if
	gender := "man"
	age2 := 18

	var canWork bool

	if gender == "man" {
		if age2 >= 18 {
			canWork = true
		} else {
			canWork = false
		}
	} else {
		if age2 >= 20 {
			canWork = true
		} else {
			canWork = false
		}
	}

	if !canWork { // ini sama artinya dengan canwork == false
		fmt.Println("Saya tidak boleh bekerja")
	} else {
		fmt.Println("Saya boleh bekerja")
	}
}
