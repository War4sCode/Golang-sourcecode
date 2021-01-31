package main

import "fmt"

//nil, yaitu data kosong
/* nil hanya bisa digunakan dibeberapa tipe data
seperti interface, map, slice, function, pointer, dan channel */

func mapbaru(tokoh string) map[string]string {
	if tokoh == "" {
		return nil
	} else {
		return map[string]string{
			"tokoh": tokoh,
		}
	}
}
func main() {
	var orang map[string]string = mapbaru("dimas")
	fmt.Println(orang)

	if orang == nil {
		fmt.Println("ini orang")
	} else {
		fmt.Println(orang)
	}

	pahlawan := mapbaru("soekarno")
	fmt.Println(pahlawan)
}
