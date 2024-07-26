package main

import "fmt"

func main() {
	x := 0
	fmt.Print("Masukan Harga : ")
	fmt.Scanln(&x)
	fmt.Println(x)
	if x >= 500001 {
		fmt.Println("Mahal")
	} else if x > 200000 {
		fmt.Println("Sedang")
	} else {
		fmt.Println("Murah")
	}
	// switch
}
