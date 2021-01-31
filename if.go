package main

import "fmt"

func main() {
	var absen = "zain"

	if absen == "nanda wido" {
		fmt.Println("Lagi main PUBG")
	} else if absen == "maulana" {
		fmt.Println("halo maulana!boleh kenalan gak?")
	} else if absen == "nanda wido" {
		fmt.Println("halo nanda!boleh kenalan gak?")
	} else {
		fmt.Println("Hai, boleh kenalan, kak?")
	}

	//kode if short statement
	if length := len(absen); length > 10 {
		fmt.Println("kebanyakan!")
	} else {
		fmt.Println("udah selesai bro:/")
	}
}
