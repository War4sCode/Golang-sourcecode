package main

import "fmt"

func main() {
	//metode slice biasa
	var anggota = [...]string{
		"maulana",
		"dimas",
		"fatih",
		"nanda",
		"bima",
		"ilham",
		"agung",
		"arif",
		"zain",
	}

	var slice1 = anggota[3:6]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//mengubah array, maka slice juga akan ikut terubah
	slice1[1] = "bima lagi ngapin?"
	fmt.Println(anggota)

	//metode append
	var slice2 = anggota[0:]

	var slice3 = append(slice2, "andi")
	fmt.Println(slice3)
	//metode slice namun merubah array yang sudah ada
	slice3[2] = "ronald"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(anggota)

	//metode 'make slice'
	makeslice := make([]string, 3, 5)

	makeslice[0] = "maulana"
	makeslice[1] = "zulkifli"
	makeslice[2] = "nugroho"

	fmt.Println(makeslice)
	fmt.Println(len(makeslice))
	fmt.Println(cap(makeslice))

	//metode 'copy slice'
	kopislice := make([]string, len(makeslice), cap(makeslice))
	copy(kopislice, makeslice)
	fmt.Println(kopislice)

	//perbedaan antara array dan slice
	//ini array
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//ini slice
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(array)
	fmt.Println(slice)

}
