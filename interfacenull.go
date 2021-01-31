package main

import "fmt"

//interface kosong adalah interface yang tidak memiliki method apapun

func ups(angka int) interface{} {
	if angka == 1 {
		return 1
	} else if angka == 3 {
		return false
	} else {
		return "cukup"
	}
}

func main() {
	var null interface{} = ups(3)
	fmt.Println(null)
}
