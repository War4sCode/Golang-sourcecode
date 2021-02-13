package matematika

import "fmt"

func main{
	var (
		luas float32
		alas float32
		tinggi float32
	)
	fmt.Printf("masukkan alas: ")
	fmt.Scanln(&alas)
	fmt.Printf("masukkan tinggi: ")
	fmt.Scanln(&tinggi)
	luas = 0.5*alas*tinggi
	fmt.Printf("maka, luas dari segitiga siku-siku adalah ",luas)
	
}
