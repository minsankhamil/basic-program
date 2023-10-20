package main

import "fmt"

func main() {

	var n uint64

	fmt.Print("Selamat datang di program untuk menentukan bilangan kelipatan 7. \n \n")
	fmt.Print("Silahkan masukkan bilangan bulat: \n")
	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Print("Masukkan tidak valid")
		return
	}

	var bilkel7 uint64 = n % 7

	if bilkel7 == 0 {
		fmt.Print("\nBilangan ", n, " termasuk bilangan kelipatan 7")
	} else {
		fmt.Print("\nBilangan ", n, " bukan termasuk bilangan kelipatan 7")
	}
	//selesai
}
