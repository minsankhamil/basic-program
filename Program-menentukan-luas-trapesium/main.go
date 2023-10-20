package main

import (
	"fmt"
)

func main() {

	var a float32
	var b float32
	var h float32

	fmt.Print("Selamat datang di program untuk menentukan luas trapesium. \n \n")
	fmt.Print("Silahkan masukkan nilai panjang sisi atas trapesium: \n")
	_, erra := fmt.Scan(&a)

	if erra != nil {
		fmt.Print("Masukkan tidak valid")
		return
	}

	fmt.Print("Silahkan masukkan nilai panjang sisi bawah trapesium: \n")
	_, errb := fmt.Scan(&b)

	if errb != nil {
		fmt.Print("Masukkan tidak valid")
		return
	}

	fmt.Print("Silahkan masukkan nilai tinggi trapesium: \n")
	_, errh := fmt.Scan(&h)

	if errh != nil {
		fmt.Print("Masukkan tidak valid")
		return
	}

	if a >= 0 && b >= 0 && h >= 0 {
		/*fungsi untuk menghitung luas trapesium*/
		var luasTrapesium float32 = ((a + b) / 2) * h

		fmt.Print("Maka luas trapesium adalah: \n")
		fmt.Print(luasTrapesium)
	} else {
		fmt.Println("Luas trapesium tidak dapat ditentukan karena terdapat input yang bukan merupakan bilangan bulat positif.")
	}
	//selesai
}
