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
	var name string
	fmt.Println("\n ini adalah perhitungan energi potensial")
	fmt.Printf("siap lanjut? ")
	fmt.Scan(&name)
	fmt.Printf("ok\n\n")
	sainsdanteknologi.EnergiPotensial()
	var nama string
	fmt.Println("\n ini adalah perhitungan energi kinetik")
	fmt.Printf("siap lanjut? ")
	fmt.Scan(&nama)
	fmt.Printf("ok\n\n")
	sainsdanteknologi.EnergiKinetik()
	var jeneng string
	fmt.Println("\n ini adalah perhitungan getaran")
	fmt.Printf("siap lanjut? ")
	fmt.Scan(&jeneng)
	fmt.Printf("ok\n\n")
	sainsdanteknologi.Frekuensi()
}

