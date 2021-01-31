package main

import "fmt"

func main() {
	nama := "maulana"
	hitung := 0

	menambah := func() {
		nama := "nanda"
		fmt.Println("increment")
		hitung++
		fmt.Println(nama)
	}

	menambah()
	fmt.Println(hitung)
	fmt.Println(nama)
	fmt.Println(nama)
}
