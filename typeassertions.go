package main

import "fmt"

//untuk merubah tipe data menjadi tipe data yang diinginkan
func random() interface{} {
	return "maulana"
}

func main() {
	//penggunaan switch pada type assertion lebih mempan daripada membuat manual
	var resultz interface{} = random()
	switch value := resultz.(type) {
	case string:
		fmt.Println("value", value, "is string")
	case int:
		fmt.Println("value", value, "is integer")
	case bool:
		fmt.Println("value", value, "is boolean")
	default:
		fmt.Println("tidak dikenali")
	}
}
