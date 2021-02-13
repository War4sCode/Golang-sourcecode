package matematika

import (
	"fmt"
	"math"
)

func Lingkaran(){
	var (
		phi float64 = 3.14
		r float64
		luas float64
	)
	
	fmt.Printf("Masukkan jari-jari: ")
	fmt.Scan(&r)
	luas = phi*(math.Pow(r,2))
	fmt.Println(luas)
}

