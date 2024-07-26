package main

import "fmt"

type mahasiswa struct {
	Nama, Npm, Jurusan, Fakultas string
}

func main() {
	// fadli
	// yoga
	// bagas
	// salman

	var fadli mahasiswa
	fadli.Nama = "Achmad Fadli Iskandar"
	fadli.Npm = "50422069"
	fadli.Jurusan = "informatika"
	fadli.Fakultas = "teknologi industri"

	var yoga mahasiswa
	yoga.Nama = "yoga windy fajar"
	yoga.Npm = "51422069"
	yoga.Jurusan = "informatika"
	yoga.Fakultas = "teknologi industri"

	var bagas mahasiswa
	bagas.Nama = "ramadya bagas"
	bagas.Npm = "5021543433"
	bagas.Jurusan = "sistem informasi"
	bagas.Fakultas = "fikom"

	var salman mahasiswa
	salman.Nama = "salman alfarizi taha"
	salman.Npm = "50312323232"
	salman.Jurusan = "teknik elektro"
	salman.Fakultas = "teknologi industri"

	fmt.Println(fadli)
	fmt.Println(yoga)
	fmt.Println(bagas)
	fmt.Println(salman)
}
