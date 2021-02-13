package sainsdanteknologi

import "fmt"

func EnergiKinetik(){
	var (
		massa float32
		kecepatanbenda float32
		ekinetik float32
	)
	fmt.Printf("massanya berapa? ")
	fmt.Scan(&massa)
	fmt.Printf("kecepatannya berapa? ")
	fmt.Scan(&kecepatanbenda)
	ekinetik = massa*kecepatanbenda*kecepatanbenda*0.5
	fmt.Printf("maka ennergi kinetiknya adalah: ",ekinetik)
}