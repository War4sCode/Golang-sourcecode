package main

import "fmt"

type callname interface {
	menyapa() string
}

func sayHi(sapa callname) {
	fmt.Println("halo", sapa.menyapa())
}

func caution(peringatan callname) {
	fmt.Println("awas ada", peringatan.menyapa())
}

func Run(lari callname) {
	fmt.Println("lari woy,", lari.menyapa())
}

type siswa struct {
	nama string
}

type hewan struct {
	jenis string
}

type run struct {
	kalimat string
}

func (Siswa siswa) menyapa() string {
	return Siswa.nama
}

func (Hewan hewan) menyapa() string {
	return Hewan.jenis
}

func (Lari run) menyapa() string {
	return Lari.kalimat
}

func main() {
	var fatih siswa
	fatih.nama = "Dimas"
	sayHi(fatih)

	var animal hewan
	animal.jenis = "anjing"
	caution(animal)

	var Larirun run
	Larirun.kalimat = "jangan diem-diem tok!"
	Run(Larirun)
}
