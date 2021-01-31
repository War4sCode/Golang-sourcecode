package main

import "fmt"

//struct adalah kumpulan dari field
type datadiri struct {
	nama, alamat, sekolah string
	absen, NIS            int
}

func (Customer datadiri) sayhello(nama string) {
	fmt.Println("halo ", nama, "nama saya ", Customer.nama)
}

func main() {
	var maul datadiri
	maul.nama = "maulana zulkifli nugroho"
	maul.alamat = "Boyolali"
	maul.sekolah = "Barkab"
	maul.absen = 17
	maul.NIS = 4563

	fmt.Println(maul.nama)
	fmt.Println(maul.alamat)
	fmt.Println(maul.sekolah)
	fmt.Println(maul.absen)
	fmt.Println(maul.NIS)

	//struct method, menyambung di baris ke 11
	maul.sayhello("wawan")

	//struct literals metode dengan lebih sederhana
	fatih := datadiri{
		nama:    "fatih kurniawan",
		alamat:  "Boyolali",
		sekolah: "barkab",
		absen:   10,
		NIS:     4325,
	}
	fmt.Println(fatih)
}
