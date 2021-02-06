package matematika

import "fmt" 

func PersegiPanjang(){
	var (
		panjang int
		lebar int
		luas int
	)
	var penyebut int = 2 
	fmt.Printf("masukkan panjangnya: ")
	fmt.Scan(&panjang)
	fmt.Printf("masukkan lebarnya: ")
	fmt.Scan(&lebar)
	luas = (penyebut) * (panjang + lebar )
	fmt.Printf("jadi luas dari persegi  panjang adalah: ",luas)
}
