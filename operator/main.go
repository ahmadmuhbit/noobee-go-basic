package main

import "fmt"

func main() {
	// Operator aritmatika
	fmt.Println(1 + 1)
	fmt.Println(2 - 1)
	fmt.Println(2 * 5)
	fmt.Println(4 / 2)
	fmt.Println(5 % 2)

	var num1 int = 10
	var num2 int = 2
	str1 := "Noo"
	str2 := "Bee"

	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)
	fmt.Println(str1 + str2)

	// Operator increment dan decrement
	i := 5
	i++ // i = i + 1

	j := 5
	j-- // i = i - 1

	fmt.Println(i)
	fmt.Println(j)

	// Operator penugasan
	var a, b, c, d, e = 5, 10, 15, 20, 25

	var x = 5
	a += 5 // a = a + 5
	b -= 5 // a = a - 5
	c *= 5 // a = a * 5
	d /= 5 // a = a / 5
	e %= 5 // a = a % 5

	fmt.Println("nilai x:", x)
	fmt.Println("nilai a:", a)
	fmt.Println("nilai b:", b)
	fmt.Println("nilai c:", c)
	fmt.Println("nilai d:", d)
	fmt.Println("nilai e:", e)
}
