package main

import "fmt"

func main() {
	var (
		nama1 = "fatih"
		nama2 = "wido"
	)
	var (
		hasil  bool = nama1 == nama2
		hasil1 bool = nama1 <= nama2
	)
	fmt.Println(hasil)
	fmt.Println(hasil1)

	var nilai = 90
	var nilai1 = 50

	fmt.Println(nilai == nilai1)
	fmt.Println(nilai >= nilai1)
	fmt.Println(nilai <= nilai1)
	fmt.Println(nilai != nilai1)

}
