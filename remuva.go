package main

import "fmt"

func hobi() (string, string, string) {
	return "nanda wido", "suka main PUBG", "dan buat video"
}

func dolan() (string, string, string, string) {
	return "Ronaldjidin", "suka nonton film", "dan", "main free fire"
}

func main() {
	satu, dua, tiga := hobi()
	fmt.Println(satu, dua, tiga)
	//menghiraukan salah satu value, dengan cara memberikan tanda garis bawah
	siji, loro, _, _ := dolan()
	fmt.Println(siji, loro)
}
