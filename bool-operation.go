package main

import "fmt"

func main() {
	//cara lengkapnya
	var nilaiUN = 70
	var nilaiUKK = 80

	var lulus1 bool = nilaiUN <= 70
	var lulus2 bool = nilaiUKK >= 80
	fmt.Println(lulus1)
	fmt.Println(lulus2)

	var result bool = lulus1 || lulus2

	fmt.Println(result)

	//cara singkat
	fmt.Println(nilaiUN >= 70 && nilaiUKK >= 80)
}
