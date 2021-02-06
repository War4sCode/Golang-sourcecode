package matematika

import "fmt"

func SikuKeliling(){
	var (
		sisi1 int
		sisi2 int
		sisi3 int
		hasil int
	)
	fmt.Printf("masukkan sisi pertama: ")
	fmt.Scan(&sisi1)
	fmt.Printf("masukkan sisi kedua: ")
	fmt.Scan(&sisi2)
	fmt.Printf("masukkan sisi ketiga: ")
	fmt.Scan(&sisi3)
	
	hasil = 1*sisi1+sisi2+sisi3
	
	fmt.Printf("maka hasil keliling dari segitiga siku-siku adalah: ", hasil)
	
}

