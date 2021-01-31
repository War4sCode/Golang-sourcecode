package main

import "fmt"

func main() {
	type NoKtp string
	type GraduateStatus bool

	var Nomor NoKtp = "15433675342"
	var lulusSMK GraduateStatus = false
	fmt.Println(Nomor)
	fmt.Println(lulusSMK)
}
