package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Konversi tipe data integer
	var a int8 = 10

	var b int16 = int16(a) // Konversi int8 ke int16
	fmt.Printf("Tipe data a: %T, nilai a: %v\n", a, a)
	fmt.Printf("Tipe data b: %T, nilai b: %v\n", b, b)

	// Konversi tipe data int ke float, dan float ke int
	var c float32 = 20.5

	var d float64 = float64(c) // Konversi float32 ke float64
	var e float64 = float64(a) // Konversi int ke float
	var f int = int(c)         // Konversi float ke int

	fmt.Printf("Tipe data d: %T, nilai d: %v\n", d, d)
	fmt.Printf("Tipe data e: %T, nilai e: %v\n", e, e)
	fmt.Printf("Tipe data f: %T, nilai f: %v\n", f, f)

	// Konversi dari byte ke string dan string ke byte
	str1 := "Haii"
	byte1 := []byte(str1) // Konversi dari string ke byte
	fmt.Printf("str1: %v\n", str1)
	fmt.Printf("byte1: %v\n", byte1)

	fmt.Printf("Cast dari byte1 ke string: %v\n", string(byte1)) // Konversi dari byte ke string
	fmt.Printf("Cast dari byte ke string: %v\n", string(65))

	// Konversi dari integer ke string
	var num1 int = 1
	strNum := strconv.Itoa(num1)
	fmt.Printf("Tipe data: %T, nilai: %v\n", strNum, strNum)

	// Konversi dari string ke integer
	var str2 string = "1"

	numStr2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Tipe data: %T, nilai: %v\n", numStr2, numStr2)
	}

	// Konversi dari string ke integer parseint
	var str3 string = "1234"

	numStr3, err := strconv.ParseInt(str3, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(numStr3)
	}

	// Konversi string ke floating point
	numFloat, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Tipe data: %T, nilai: %v\n", numFloat, numFloat)
	}

	// Konversi string ke bool
	bString := "true"

	boolean, err := strconv.ParseBool(bString)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Tipe data: %T, nilai: %v\n", boolean, boolean)
	}

	// Konversi bool ke string
	boolean2 := true

	booleanString := strconv.FormatBool(boolean2)
	fmt.Printf("Tipe data: %T, nilai: %v\n", booleanString, booleanString)
}
