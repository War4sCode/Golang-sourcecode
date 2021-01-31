package main

import "fmt"

//function named return values ini hanya ada di golang,
func getGame() (suka string, ganti string, sambung string, nomor int, akhir string) {
	suka = "dimas"
	ganti = "suka main CODM"
	sambung = "sejak"
	nomor = 200
	akhir = "hari yang lalu"
	return
}

func main() {
	suka, ganti, sambung, nomor, akhir := getGame()
	fmt.Println(suka, ganti, sambung, nomor, akhir)
}
