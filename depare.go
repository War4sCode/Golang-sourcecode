package main

import "fmt"

/* defer adalah function yang bisa dijadwalkan untuk dieksekusi
setelah sebuah function selesai dieksekusi*/
//jika tidak memakai defer, maka program akan dieksekusi sesuai kode bloknya

func login() {
	fmt.Println("Selamat datang di golang!")
}

func Run() {
	defer login()
	fmt.Println("halo, maulana")
}

func main() {
	Run()
	endrun(true)
}

/* panic adalah function yang digunakan untuk menghentikan program
panic dipanggil saat program yang dibuat, ada yang error
saat panic dipanggil, program akan berhenti, namun function defer akan tetap dieksekusi */
func runapp() {
	//contoh penggunaan function recover
	pesan := recover()
	fmt.Println("error dengan pesan", pesan)
	fmt.Println("selesai")
}

func endrun(eror bool) {
	defer runapp()
	if eror {
		//contoh penggunaan function panic
		panic("eror bro:/")
	}
	fmt.Println("aplikasi masih dalam debugging")
}

/*recover adalah function yang digunakan untuk menangkap data panic
dengan recover panic akan terhenti, sehingga program akan tetap berjalan*/
