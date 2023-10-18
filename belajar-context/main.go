package main

import (
	"context"
	"fmt"
)

func main() {
	// Inisialisasi context
	// context := context.Background()
	// fmt.Println(context)

	// Context with value
	ctx := context.Background()

	ctxLevel1_A := context.WithValue(ctx, "a", "A")
	ctxLevel1_B := context.WithValue(ctx, "b", "B")

	ctxLevel2_A1 := context.WithValue(ctxLevel1_A, "a1", "A1")
	ctxLevel2_A2 := context.WithValue(ctxLevel1_A, "a2", "A2")
	ctxLevel2_A3 := context.WithValue(ctxLevel1_A, "a3", "A3")

	ctxLevel2_B1 := context.WithValue(ctxLevel1_B, "b1", "B1")
	ctxLevel2_B2 := context.WithValue(ctxLevel1_B, "b2", "B2")

	fmt.Println(ctx)
	fmt.Println(ctxLevel1_A)
	fmt.Println(ctxLevel1_B)

	fmt.Println()

	fmt.Println(ctxLevel2_A1)
	fmt.Println(ctxLevel2_A2)
	fmt.Println(ctxLevel2_A3)

	fmt.Println()

	fmt.Println(ctxLevel2_B1)
	fmt.Println(ctxLevel2_B2)

	// Mengambil value
	fmt.Println(ctxLevel2_B1.Value("b1"))
	fmt.Println(ctxLevel2_B1.Value("b2"))
	fmt.Println(ctxLevel2_B1.Value("a3"))

	fmt.Println(ctxLevel2_B2.Value("b"))
}
