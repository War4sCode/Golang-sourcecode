package main

import "fmt"

//simbol '*' adalah pointer
type alamat struct {
	Kota, Provinsi, Negara string
}

func merubahkebelanda(rumah *alamat){
	rumah.Negara = "Netherland"
}

func main() {
	Tempattinggal := alamat{"Boyolali", "jawa tengah", "indonesia"}
	//simbol &, menandakan bukan menjadi sendiri lagi, namun mengikuti yang pointernya
	Tempatsinggah := &Tempattinggal

	Tempatsinggah.Kota = "surakarta"

	*Tempatsinggah = alamat{"malang", "jawa timur", "indonesia"}

	fmt.Println(Tempatsinggah)
	fmt.Println(Tempattinggal)

	var address *alamat = new(alamat)
	fmt.Println(address)
	
	var home = alamat {
		Kota: "Boyolali",
		Provinsi: "jawa tengah",
		Negara: "",
	}
	
	var homepointer *alamat = &home
	merubahkebelanda(homepointer)
	fmt.Println(home)
}
