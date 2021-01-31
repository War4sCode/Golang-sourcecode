package main

import "fmt"

func getnama(absen string) string {
	if absen == "17" {
		return "absen saya adalah " + absen
	} else {
		return "salah, absenmu bukan " + absen
	}
}

func main() {
	hasil1 := getnama("17")
	fmt.Println(hasil1)
}
