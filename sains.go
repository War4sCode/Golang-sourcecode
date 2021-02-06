package main

import (
	"golang-sourcecode/sainsdanteknologi"
	"fmt"
)

func main(){
	var orang string
	fmt.Printf("siapa namamu?")
	fmt.Scan(&orang)
	fmt.Printf(" ok, selamat datang di programku,kak" + orang)
	
	var science int
	fmt.Println("Apa yang igin dicari?(ketikkan angkanya saja)")
	fmt.Println("/t1. daya")
	fmt.Println("/t2. usaha")
	fmt.Println("/t3. waktu")
	fmt.Println("/t4. Gerak lurus beraturan-kecepatan")
	fmt.Println("/t5. Energi Potensial")
	fmt.Println("/t6. Energi Kinetik")
	fmt.Println("/t7. Frekuensi")
	fmt.Printf("masukkan angka yang dipilih")
	fmt.Scan(&science)
	
	switch science{
		case 1: sainsdanteknologi.Daya()
		case 2: sainsdanteknologi.Usaha()
		case 3: sainsdanteknologi.Waktu()
		case 4: sainsdanteknologi.GLB_Fast()
		case 5: sainsdanteknologi.EnergiPotensial()
		case 6: sainsdanteknologi.EnergiKinetik()
		case 7: sainsdanteknologi.Frekuensi()
		default: sainsdanteknologi.Daya()
	}
}

