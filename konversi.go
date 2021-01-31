package main

import "fmt"

func main() {
	var nilai64 int64 = 1200
	var nilai32 int32 = int32(nilai64)
	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai16)

	fmt.Println(nilai64)
	fmt.Println(nilai32)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	var nama = "maulana zulkifli Nugroho"
	var z = nama[1]
	var zString string = string(z)

	fmt.Println(nama)
	fmt.Println(zString)
}
