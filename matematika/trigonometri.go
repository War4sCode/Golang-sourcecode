package matematika

import (
	"fmt"
	"math"
)

func Trigonometri(){
	var sin, nilai float64
	fmt.Println("Masukkan nilai sinus: ")
	fmt.Scanf("%f", &sin)
	nilai = math.Sin(sin)
	fmt.Println(nilai)
	
	var cos, nilai2 float64
	fmt.Println("Masukkan nilai cosan: ")
	fmt.Scanf("%f", &cos)
	nilai2 = math.Cos(cos)
	fmt.Println(nilai2)
	
	var tan, nilai3 float64
	fmt.Println("Masukkan nilai tangen: ")
	fmt.Scanf("%f", &tan)
	nilai3 = math.Tan(tan)
	fmt.Println(nilai3)
}