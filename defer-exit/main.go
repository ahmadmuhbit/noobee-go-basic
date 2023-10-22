package main

import (
	"fmt"
	"os"
)

// Defer
func cetak(text string) {
	fmt.Println(text)
}

func main() {
	// // Defer
	// defer cetak("Text 1")
	// defer cetak("Text 2")
	// cetak("Text 3")
	// defer cetak("Text 4")
	// cetak("Text 5")

	num1 := 5

	if num1%2 > 0 {
		cetak("Ini adalah bilangan ganjil")
		// defer cetak("akan berakhir setelah proses diatas selesai")
		func() {
			defer cetak("akan berakhir setelah proses diatas selesai")
		}()
	}

	cetak("Oh tentu tidak, Defer di eksekusi seelah ini")

	// Exit
	names := []string{"NooB", "Bee", "Go", "NodeJS"}

	for _, name := range names {
		if name == "Go" {
			os.Exit(1)
		}
		cetak(name)
	}
}
