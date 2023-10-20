package main

import (
	"fmt"
)

func main() {

	var a int

	fmt.Print("Selamat datang di program untuk menentukan bilangan prima. \n \n")
	fmt.Println("Silahkan masukkan bilangan bulat:")
	_, erra := fmt.Scan(&a)

	if erra != nil {
		fmt.Scan("Masukkan tidak valid")
	}

	if a < 2 {
		fmt.Print("Bilangan ", a, " bukan merupakan bilangan prima")
	} else {
		for n := 2; n != a; n++ {
			if a%n == 0 {
				fmt.Println("Bilangan ", a, " bukan merupakan bilangan prima")
				return
			}
		}
		fmt.Println("Bilangan ", a, " merupakan bilangan prima")
	}
}
