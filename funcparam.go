package main

import "fmt"

func sayhobby1(hobisatu string, hobidua string) {
	fmt.Println("hobiku adalah ", hobisatu, " dan", hobidua)
}

func sayhobby2(tiga string, cod string) {
	fmt.Println("hobinya ", tiga, " adalah ", cod)
}

func main() {
	sayhobby1("bersepeda", "basket")
	sayhobby2("dimas", "main CODM")
}
