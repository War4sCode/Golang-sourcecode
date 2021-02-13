package sainsdanteknologi

import "fmt"

func Frekuensi(){
	var (
		getaran int
		waktu int
		frekuensi int
	)
	
	fmt.Printf("jumlah getarannya berapa? ")
	fmt.Scan(&getaran)
	fmt.Printf("lama waktunya berapa? ")
	fmt.Scan(&waktu)
	frekuensi = getaran/waktu
	fmt.Printf("maka Frekuensinya adalah: ",frekuensi)
}