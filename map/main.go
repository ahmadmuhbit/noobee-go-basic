package main

import "fmt"

func main() {
	// Membuat map
	var person = map[string]string{
		"name": "NooBee",
		"job":  "Programmer",
	}

	fmt.Println(person)

	computer := map[string]int{
		"stock": 50,
		"harga": 20000000,
	}
	fmt.Println(computer)

	// Membuat map dengan function make
	laptop := make(map[string]string)
	laptop["merk"] = "Apple"
	laptop["model"] = "Macbook Pro"

	fmt.Println(laptop)

	// Membuat map kosong
	mapKosong1 := make(map[string]string)
	fmt.Println(mapKosong1)

	// Mengakses elemen map
	fmt.Println(person["name"])
	fmt.Println(person["job"])

	// Menambah dan mengubah elemen map
	laptop["tahun"] = "2020"        // Menambah
	laptop["model"] = "Macbook Air" // Merubah

	fmt.Println(laptop)
	fmt.Println(len(laptop))
}
