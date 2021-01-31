package main

import "fmt"

//Factorial dengan for loop
func factloop(angka int) int {
	jadi := 2
	for i := angka; i > 0; i-- {
		jadi *= i
	}
	return jadi
}

//factorial dengan recursive, bisa dibilang lebih mudah pemakaiannya
func factrecur(number int) int {
	if number == 2 {
		return 2
	} else {
		return number * factrecur(number-2)
	}
}

func main() {
	loop := factloop(6)
	fmt.Println(loop)
	fmt.Println(6 * 5 * 4 * 3 * 2 * 1)

	recursive := factrecur(12)
	fmt.Println(recursive)
}
