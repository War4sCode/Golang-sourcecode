package sainsdanteknologi

import "fmt"

func EnergiPotensial(){
	var (
		massa int
		pgrav int
		epotensial int
		tinggi int
	)
	
	fmt.Printf("Berapa massanya? ")
	fmt.Scan(&massa)
	fmt.Printf("Berapa gravitasinya? ")
	fmt.Scan(&pgrav)
	fmt.Printf("Berapa tinggi bendanya? ")
	fmt.Scan(&tinggi)
	epotensial = massa*pgrav*tinggi
	fmt.Printf("maka energi potensialnya adalah: ",epotensial)
}
