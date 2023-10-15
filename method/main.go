package main

import "fmt"

type Car struct {
	Name  string
	Color string
}

func (c Car) GetName() string {
	return c.Name
}

func (c Car) SayHello() {
	fmt.Println("Hello dari", c.Name, c.Color)
}

func main() {
	car := Car{
		Name:  "Civic",
		Color: "White",
	}

	nameCar := car.GetName()
	fmt.Println("Nama mobilnya adalah:", nameCar)

	car.SayHello()
}
