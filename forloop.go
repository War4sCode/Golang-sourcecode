package main

import "fmt"

func main() {
	time := 0
	for time < 10 {
		fmt.Println("ini sudah yang ke ", time, "kalinya")
		time += 5
	}
	//init statement dan post statement
	for counter := 0; counter < 5; counter++ {
		fmt.Println("ini sudah yang ke", counter)
	}
	//for untuk mengeluarkan variabel tertentu, seperti array, slice, maupun map(metode kompleks)
	mapel := []string{"bahasa indonesia", "PKN", "bahasa inggris", "produktif"}

	for i := 0; i < len(mapel); i++ {
		fmt.Println(mapel[i])
	}
	//metode for range(metode gampang, sama kayak yang di atas)
	for i, value := range mapel {
		fmt.Println("index", i, "ke", value)
	}
}
