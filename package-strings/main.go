package main

import (
	"fmt"
	"strings"
)

func main() {
	// Fungsi strings.Contains
	contains := strings.Contains("seafood", "foo")
	fmt.Println(contains)
	fmt.Println(strings.Contains("seafood", "bar"))

	str1 := "Nama saya NooBeeID"
	cari := "NooBeeID"
	isExist := strings.Contains(str1, cari)
	fmt.Printf("Apakah string %s ada di dalam text %s ? %v\n", cari, str1, isExist)

	// Fungsi strings.Replace
	replaceString := "Nama saya NooBeeID"
	fmt.Println("Text asli:", replaceString)
	fmt.Println("Ubah 1 huruf a:", strings.Replace(replaceString, "a", "o", 1))
	fmt.Println("Ubah 2 huruf a:", strings.Replace(replaceString, "a", "o", 2))
	fmt.Println("Ubah 3 huruf a:", strings.Replace(replaceString, "a", "o", 3))
	fmt.Println("Ubah semua huruf a:", strings.Replace(replaceString, "a", "o", -1))
	fmt.Println(strings.ReplaceAll(replaceString, "a", "o"))

	// Fungsi strings.Split
	str2 := "Nama saya NooBeeID"

	splitStr := strings.Split(str2, " ")
	fmt.Println(str2)
	fmt.Println(splitStr)
	fmt.Println(splitStr[0])

	// Fungsi strings.Join
	s := []string{"Nama", "saya", "NooBeeID"}
	fmt.Println(strings.Join(s, "-"))
	fmt.Println(strings.Join(s, ", "))

	// Fungsi strings.ToUpper dan strings.ToLower
	fmt.Println(strings.ToUpper("Nama saya NooBeeID"))
	fmt.Println(strings.ToLower("Nama saya NooBeeID"))
}
