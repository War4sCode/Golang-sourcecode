package main

import "fmt"

func hitung(angka ...int) int {
	total := 0
	for _, value := range angka {
		total += value
	}

	return total
}

func main() {
	total := hitung(12, 24, 32, 54, 46)
	fmt.Println(total)
	//slice parameter
	number := []int{12, 24, 32, 54, 46}
	total = hitung(number...)
	fmt.Println(total)
}
