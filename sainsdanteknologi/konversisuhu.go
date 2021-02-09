package sainsdanteknologi

import "fmt"

func Celcius(){
	var cel float64
	
	fmt.Printf("Masukkan nilai celcius: \n")
	fmt.Scan(&cel)
	fmt.Println("Kelvin = ",cel+273.15)
	fmt.Println("Fahrenheit = ",cel*1.8+32)
	fmt.Println("Reamur = ",cel*0.8)
}

func Kelvin(){
	var kelvin float64
	
	fmt.Printf("Masukkan nilai kelvin: \n")
	fmt.Scan(&kelvin)
	fmt.Println("celcius = ",kelvin-273.15)
	fmt.Println("Fahrenheit = ",kelvin*1.8-459.67)
	fmt.Println("Reamur = ",kelvin-273.15*0.8)
}

func Fahrenheit(){
	var fahren float64
	
	fmt.Printf("Masukkan nilai fahrenheitnya: \n")
	fmt.Scan(&fahren)
	fmt.Println("celcius = ",fahren-32/1.8)
	fmt.Println("kelvin = ",fahren+459.67/1.8)
	fmt.Println("Reamur = ",fahren-32*2.25)
}

func Reamur(){
	var mur float64
	
	fmt.Printf("Masukkan nilai reamurnya: ")
	fmt.Scan(&mur)
	fmt.Println("celcius = ",mur/0.8)
	fmt.Println("kelvin = ",mur/0.8+273.15)
	fmt.Println("fahrenheit = ",mur*2.25+32)
}