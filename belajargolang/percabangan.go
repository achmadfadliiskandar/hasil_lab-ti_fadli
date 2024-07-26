package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("tugas percabangan")
	var nilai string
	var output string
	fmt.Print("Masukkan Nilai : ")
	fmt.Scanln(&nilai)
	myqin, err := strconv.Atoi(nilai)
	// fmt.Println(nilai)
	if err == nil {
		if myqin > 100 {
			output = "nilai melampaui batas"
		} else if myqin > 85 {
			output = "A"
		} else if myqin > 75 {
			output = "B"
		} else if myqin > 60 {
			output = "C"
		} else if myqin > 50 {
			output = "D"
		} else if myqin < 0 {
			output = "nilai mengurangi batas"
		} else {
			output = "E"
		}
		fmt.Println("Output : " + output)
	} else {
		fmt.Println("inputnya angka aja ya!")
	}
}
