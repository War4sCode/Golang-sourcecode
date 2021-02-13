package main

import "fmt"

func main() {
	var absen = "zain"
	
	fmt.Printf("ketik nama\n")
	fmt.Printf("Berikut nama yang tersedia:\n")
	fmt.Printf("1.zain\n")
	fmt.Printf("2.nanda wido\n")
	fmt.Printf("3.maulana\n")
	fmt.Printf("4.terserah kamu\n")
	fmt.Printf("ketik nama sesuai yang ada di list ya!\n")
	fmt.Scan(&absen)
	
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
