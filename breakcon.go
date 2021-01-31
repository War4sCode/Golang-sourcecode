package main

import "fmt"

func main() {
	//metode break
	for i := 0; i < 30; i++ {
		if i == 10 {
			break
		}
		fmt.Println("ini adalah perulangan ke ", i)
	}
	//metode continue
	for waktu := 0; waktu < 59; waktu++ {
		if waktu%2 == 1 {
			continue
		}
		fmt.Println("ini adalah perulangan ke", waktu)
	}
}
