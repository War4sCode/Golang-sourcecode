package main

import "fmt"

func sayJalan() {
	fmt.Println("siap!")
}

func saymalam() {
	fmt.Println("selamat malam!")
}

func main() {
	//output biasa tanpa perulangan
	sayJalan()
	//output dengan perulangan
	for i := 0; i < 15; i++ {
		saymalam()
	}

}
