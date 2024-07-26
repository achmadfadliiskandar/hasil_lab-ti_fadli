package main

import "fmt"

type peserta struct {
	nama     string
	npm      string
	jurusan  string
	fakultas string
}

func main() {
	x1 := 0
	x2 := 0
	peserta1 := peserta{nama: "Jaenuddin", npm: "50443322", jurusan: "informatika", fakultas: "teknologi industri"}
	peserta2 := peserta{nama: "Marsella", npm: "50554433", jurusan: "informatika", fakultas: "teknologi industri"}

	//tuliskan kodingan logika dibawah sini
	fmt.Println("Tentukan Pemenang")
	fmt.Print("Masukan Nilai p1 : ")
	fmt.Scanln(&x1)
	fmt.Print("Masukan Nilai p2 : ")
	fmt.Scanln(&x2)
	fmt.Println("Pemenangnya adalah ... ")

	if x1 > x2 {
		fmt.Println("Peserta 1")
		fmt.Println("bernama\t" + peserta1.nama)
		fmt.Println("dengan npm\t" + peserta1.npm)
		fmt.Println("dari jurusan\t" + peserta1.jurusan)
		fmt.Println("dan fakultas\t" + peserta1.fakultas)
	} else if x1 < x2 {
		fmt.Println("Peserta 2")
		fmt.Println("bernama\t" + peserta2.nama)
		fmt.Println("dengan npm\t" + peserta2.npm)
		fmt.Println("dari jurusan\t" + peserta2.jurusan)
		fmt.Println("dan fakultas\t" + peserta2.fakultas)
	} else {
		fmt.Println("Masukan Nilai p1 : ", x1)
		fmt.Println("Masukan Nilai p2 : ", x2)
		fmt.Printf("nilai seri, keduanya adalah pemenang")
	}
}
