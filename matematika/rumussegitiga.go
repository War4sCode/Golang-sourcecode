package matematika

import "fmt"

func Segitiga(){
	var (
		alas float32
		tinggi float32
		luas float32
		
	)
	
	fmt.Printf("masukkan alas: ")
	fmt.Scanln(&alas)
	fmt.Printf("masukkan tinggi: ")
	fmt.Scanln(&tinggi)
	luas = 0.5*alas*tinggi
	fmt.Printf("jadi luas dari segitiga adalah: ", luas)
}
