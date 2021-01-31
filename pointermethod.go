package main

import "fmt"

type graduatedschool struct{
	statusin string
}

func (Graduated *graduatedschool) graduate(){
	Graduated.statusin = "lulus " + Graduated.statusin
	fmt.Println("sebenarnya, memang ", Graduated.statusin)
}

func main(){
	mauls := graduatedschool{"(tapi bohong)"}
	mauls.graduate()
	
	fmt.Println(mauls.statusin)
}
