package main

import "fmt"

func main() {
	name := "hasan ali wibowo"

	switch name {
	case "maulana":
		fmt.Println("halo maulana!")
	case "wido":
		fmt.Println("halo kak!ajarin main PUBG dong!")
		fmt.Println("kalau gak mau juga gapapa kok:)")
	case "budi":
		fmt.Println("hai")
	default:
		fmt.Println("siapa?")
		fmt.Println("gak kenal")
	}
	//switch dengan short statement
	switch length := len(name); length > 4 {
	case true:
		fmt.Println("Tolonglah, kak:(")
	case false:
		fmt.Println("yaudahlah")
	}
	//switch dengan tanpa kondisi, kondisi langsung di tulis di case
	length := len(name)
	switch {
	case length > 4:
		fmt.Println("namamu pendek sekali, bang!")
	case length > 10:
		fmt.Println("namamu kepanjangan!")
	default:
		fmt.Println("benar!")
	}
}
