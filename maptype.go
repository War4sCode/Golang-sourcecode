package main

import "fmt"

func main() {
	person := map[string]string{
		"nama":   "maulana",
		"alamat": "boyolali",
	}

	//mengubah data di map dengan key
	person["role"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])
	fmt.Println(person["role"])

	//membuat map baru
	var E_book = make(map[string]string)
	E_book["judul"] = "think like a programmer"
	E_book["author"] = "V.anton spraul"
	E_book["revisi"] = "belum"
	//output sebelum dihapus
	fmt.Println(E_book)
	fmt.Println(len(E_book))

	//menghapus data yang ada di map
	delete(E_book, "revisi")

	//output setelah dihapus
	fmt.Println(E_book)
	fmt.Println(len(E_book))
	fmt.Println(E_book["judul"])
	fmt.Println(E_book["author"])
}
