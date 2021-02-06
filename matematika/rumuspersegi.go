package matematika

import "fmt"

func Persegi(){
	var (
		sisi int
		luas int
	)

	fmt.Printf("Masukkan sisi: ")
	fmt.Scan(&sisi)
	luas = sisi*sisi
	fmt.Printf("jadi luasnya adalah: %d",luas)
	fmt.Printf("/n masukkan luasnya: ")
	fmt.Scan(&luas)
	sisi = luas/sisi
	fmt.Printf("jadi panjang sisi dari persegi adalah: ",sisi)
}