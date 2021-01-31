package main

import (
	"strings"
	"fmt"
)

func main(){
	fmt.Println(strings.Repeat("ha",2),strings.Repeat("lo",2))
	z := []string{"maulana","fatih","nanda","Bima"}
	fmt.Println(strings.Join(z,"."))
	fmt.Println(strings.ToUpper("maulana zulkifli nugroho"))
	fmt.Println(strings.ToLower("dimas bagus julians valentino"))
	fmt.Println(strings.Trim("fatih kurniawati","wati"))
	fmt.Println(strings.TrimRight("maulana zulkifli anugrah","anugrah"))
	fmt.Println(strings.TrimLeft("valentino dimas bagus julians","valentino"))
	
}