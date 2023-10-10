package main

import "fmt"

func main() {
	nilai := map[string]int{
		"Agama":      80,
		"Matematika": 90,
		"Olahraga":   70,
		"Design":     80,
	}

	fmt.Println("Sebelum dihapus", nilai)
	fmt.Println("Length", len(nilai))

	// Menghapus item map
	delete(nilai, "Design")

	fmt.Println("Setelah dihapus", nilai)
	fmt.Println("Length", len(nilai))

	// Mencari item map
	key := "Agama"
	val, isExist := nilai[key]

	if isExist {
		fmt.Println(key, "is exist with value", val)
	} else {
		fmt.Println(key, "is not exist")
	}

	// Slice of map
	students := []map[string]string{
		map[string]string{"name": "NooBee", "major": "Teknik Komputer"},
		map[string]string{"name": "Ahmad", "major": "Sistem Informasi"},
		map[string]string{"name": "Muhbit", "major": "Teknik Informatika"},
	}

	for _, student := range students {
		fmt.Println(student["name"], "berada di jurusan", student["major"])
	}
}
