package main

import "fmt"

// Multiple return value
// func calculate(sisi float32) (float32, float32) {
func calculate(sisi float32) (luas float32, keliling float32) {
	// luas := sisi * sisi
	// keliling := sisi * 4
	// return luas, keliling

	luas = sisi * sisi
	keliling = sisi * 4
	return
}

// fungsi variadic
func sum2(nums ...int) (total int) {
	fmt.Println(nums)

	for _, num := range nums {
		total += num // total = total + num
	}

	return total
}

// Fungsi variadic contoh 2
const (
	Plus  string = "+"
	Minus string = "-"
)

func calc(operator string, numbers ...float32) (result float32) {
	for _, num := range numbers {
		if operator == Plus {
			result += num
		} else if operator == Minus {
			result -= num
		}
	}
	return result
}

func main() {
	// Multiple return value
	var sisi float32 = 4

	luas, _ := calculate(sisi)
	fmt.Println("luas dan keliling dari persegi dengan panjang sisi", sisi)
	fmt.Println("Luas", luas)
	// fmt.Println("Keliling", keliling)

	// Function closure
	sum := func(num1, num2 int) int {
		return num1 + num2
	}
	fmt.Println(sum(1, 3))

	// Contoh closure sorting
	sorting := func(arr []int) []int {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr); j++ {
				if arr[i] < arr[j] {
					temp := arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}
			}
		}
		return arr
	}
	var arr = []int{3, 9, 5, 1, 10, 4, 2}
	fmt.Println("Awal:", arr)
	fmt.Println("Setelah diurut:", sorting(arr))

	// Fungsi variadic
	fmt.Println(sum2(1, 8, 29, 3, 4, 5, 2, 9, 7))

	// variadic function slice parameter
	nums := []int{1, 8, 29, 3, 4, 5, 2, 9, 7}
	total := sum2(nums...)
	fmt.Println(total)

	// Fungsi variadic contoh 2
	nums2 := []float32{1, 4, 2, 9}
	result := calc(Plus, nums2...)
	fmt.Println(result)
}
