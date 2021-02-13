package main

import (
	"fmt"
	"golang-sourcecode/matematika"
)

func main(){
	var name string
	fmt.Println("\n ini adalah perhitungan persegi panjang")
	fmt.Printf("siap lanjut? ")
	fmt.Scan(&name)
	matematika.PersegiPanjang()
	var nama string
	fmt.Println("\n ini adalah perhitungan segitiga")
	fmt.Printf("siap lanjut? ")
	fmt.Scan(&nama)
	matematika.Segitiga()
	var namas string
	fmt.Println("\n ini adalah perhitungan segitiga siku-siku")
	fmt.Printf("siap lanjut? ")
	fmt.Scan(&namas)
	matematika.Siku()
	var names string
	fmt.Println("\n ini adalah perhitungan segitiga siku-siku dalam keliling")
	fmt.Printf("siap lanjut? ")
	fmt.Scan(&names)
	matematika.SikuKeliling()
	var namaku string
	fmt.Println("\n ini adalah perhitungan trigonometri")
	fmt.Printf("siap lanjut? ")
	fmt.Scan(&namaku)
	matematika.Trigonometri()
	var namamu string
	fmt.Println("\n ini adalah perhitungan lingkaran")
	fmt.Printf("siap lanjut? ")
	fmt.Scan(&namamu)
	matematika.Lingkaran()
	var namanda string
	fmt.Println("\n ini adalah perhitungan persegi")
	fmt.Printf("siap lanjut? ")
	fmt.Scan(&namanda)
	matematika.Persegi()
}