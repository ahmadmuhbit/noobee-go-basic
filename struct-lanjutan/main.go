package main

import "fmt"

type Parent struct {
	Nama string
	Umur int
}

type Student struct {
	OrangTua Parent // Embedded struct
	Nama     string
	Umur     int
	Kelas    string
}

// Deklarasi slice of struct
type Employee struct {
	Name        string
	Departement string
}

type Employees []Employee

func main() {
	// Penerapan embedded struct
	parent1 := Parent{
		Nama: "Andi",
		Umur: 50,
	}

	student1 := Student{
		OrangTua: parent1,
		Nama:     "Budi",
		Umur:     11,
		Kelas:    "6A",
	}

	fmt.Println(student1)
	fmt.Printf("Orang tua %s Bernama %s\n", student1.Nama, student1.OrangTua.Nama)

	student2 := Student{
		OrangTua: Parent{
			Nama: "Jojo",
			Umur: 51,
		},
		Nama:  "Bilqis",
		Umur:  10,
		Kelas: "5B",
	}

	fmt.Println(student2)
	fmt.Printf("Orang tua %s Bernama %s\n", student2.Nama, student2.OrangTua.Nama)

	// Penerapan slice of struct
	var employees Employees
	fmt.Println(employees)

	var emp1 = Employee{Name: "Emp1", Departement: "Tech"}
	employees = append(employees, emp1)
	fmt.Println(employees)

	employees = append(employees, Employee{Name: "Emp2", Departement: "Finance"})
	fmt.Println(employees)
}
