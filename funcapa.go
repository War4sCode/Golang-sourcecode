package main

import "fmt"

//function as parameter adalah salah satu komponen unik yang dimiliki golang saja
//dengan menggunakan type declaration, penulisan function bisa lebih singkat
//bisa dibilang, type declaration ini adalah alias function, sehingga bisa lebih leluasa
type Filter func(string) string

func katakotor(nama string, filter Filter) {
	sayfilter := filter(nama)
	fmt.Println("jangan ngomong" + sayfilter)
}

func spam(nama string) string {
	if nama == "bangsat" {
		return " ..."
	} else if nama == "anjing" {
		return " ..."
	} else {
		return nama
	}
}

func main() {
	katakotor("bangsat", spam)
	katakotor("anjing", spam)
	katakotor(" jancok", spam)
}
