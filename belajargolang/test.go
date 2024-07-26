package main

import "fmt"

type pegawai struct {
	Nama, Jabatan string
	gaji          int
}

func main() {
	var achmad pegawai
	achmad.Nama = "fadli"
	achmad.Jabatan = "it support"
	achmad.gaji = 200000

	var iskandar pegawai
	iskandar.Nama = "Barik"
	iskandar.Jabatan = "programmer"
	iskandar.gaji = 500000

	fmt.Println(achmad)
	fmt.Println(iskandar)
}
