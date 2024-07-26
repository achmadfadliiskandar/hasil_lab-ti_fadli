package main

import "fmt"

type barang struct {
	Nama, jenis string
	harga       int
}
type customer struct {
	Nama, alamat string
	total        int
}

func main() {
	var buku barang
	buku.Nama = "buku informatika"
	buku.jenis = "alat tulis"
	buku.harga = 20000

	var rada customer
	rada.Nama = "Fadli"
	rada.alamat = "beji depok"
	rada.total = 50000

	iskandar := barang{
		Nama:  "achmad",
		jenis: "alat tulis",
		harga: 5000,
	}
	fmt.Println(buku)
	fmt.Println(rada)
	fmt.Println(iskandar)
}
