package sainsdanteknologi

import "fmt"

func Daya(){
	var w, t float64
	
	fmt.Printf("Masukkan nilai usaha dalam satuan joule: ")
	fmt.Scan(&w)
	fmt.Printf("Masukkan nilai waktu satuan second: ")
	fmt.Scan(&t)
	fmt.Printf("jadi hasil dari daya adalah: ",w*t)
}

func Usaha(){
	var w, t float64

	fmt.Printf("Masukkan nilai daya dalam satuan watt: ")
	fmt.Scan(&w)
	fmt.Printf("Masukkan nilai waktu satuan second: ")
	fmt.Scan(&t)
	fmt.Printf("jadi hasil dari usaha adalah: ",w/t)
}

func Waktu(){
	var w, t float64
	
	fmt.Printf("Masukkan nilai usaha dalam satuan joule: ")
	fmt.Scan(&w)
	fmt.Printf("Masukkan nilai daya satuan watt: ")
	fmt.Scan(&t)
	fmt.Printf("jadi hasil dari waktu adalah: ",w/t)
}
