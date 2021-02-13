package sainsdanteknologi

import "fmt"

func GLB_Fast(){
	var waktu, jarak float64
	
	fmt.Printf("Masukkan nilai waktu dalam satuan detik: ")
	fmt.Scan(&waktu)
	fmt.Printf("Masukkan nilai jarak perpindahan dalam satuan meter: ")
	fmt.Scan(&jarak)
	fmt.Printf("kecepatan yang didapat adalah :", jarak/waktu)
}