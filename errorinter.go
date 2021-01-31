package main

import (
	"errors"
	"fmt"
)

func perkalian(nilai int, kali int) (int, error) {
	if kali == 0 {
		return 0, errors.New("angka tidak boleh nol")
	} else {
		hasil := nilai * kali
		return hasil, nil
	}
}

func main() {
	var eror error = errors.New("aduh, error")
	fmt.Println(eror.Error())

	hasil, err := perkalian(19, 0)
	if err == nil {
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("noob!, ", err.Error())
	}
}
