package main

import "fmt"

//membuat sebuah function di dalam sebuah variabel

type Blocked func(string) bool

func anom(nama string, block Blocked) {
	if block(nama) {
		fmt.Println("Youre got blocked", nama)
	} else {
		fmt.Println("Tapi boong!wkkwk, kasihan si", nama)
	}
}

func main() {
	block := func(nama string) bool {
		return nama == "bot"
	}
	anom("bot", block)
	anom("admin", block)
}
