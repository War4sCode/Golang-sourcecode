package main

import "fmt"

//kode program yang biasa atau yang sederhana
func getName(nama string) string {
	return "selamat pagi" + nama
}

func getnama(absen string) string {
	if absen == "17" {
		return "absen saya adalah " + absen
	} else {
		return "salah, absenmu bukan " + absen
	}
}

func main() {
	hasil := getName(" maulana")
	fmt.Println(hasil)
	hasil1 := getnama("17")
	fmt.Println(hasil1)
}
