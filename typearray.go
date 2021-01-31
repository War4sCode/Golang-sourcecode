package main

import "fmt"

func main() {
	var siswa [5]string

	siswa[0] = "bima"
	siswa[1] = "ronald"
	siswa[2] = "agung"
	siswa[3] = "arif"
	siswa[4] = "bagas"

	fmt.Println(siswa[0])
	fmt.Println(len("bima"))
	fmt.Println(siswa[1])
	fmt.Println(siswa[2])
	fmt.Println(siswa[3])
	fmt.Println(siswa[4])

	//secara langsung namun integer
	var siswa1 = [5]int{
		90,
		80,
		70,
		100,
		80,
	}
	fmt.Println(siswa1)

	fmt.Println(len(siswa))
	fmt.Println(len(siswa1))
}
