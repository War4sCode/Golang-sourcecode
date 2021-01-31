package main

import "fmt"

func sepeda(jenis string) string {
	return "hobinya wawan adalah " + jenis
}

func main() {
	hobi := sepeda
	jadi := hobi("bersepeda")
	fmt.Println(jadi)
	fmt.Println(sepeda("Naik becak"))
}
