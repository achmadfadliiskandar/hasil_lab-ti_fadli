package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Buku struct {
	Nomor     int    `json:"nomor"`
	Nama      string `json:"nama"`
	Pengarang string `json:"pengarang"`
}

func NewBuku() []Buku {
	bk := []Buku{
		//masukkan 5 data didalam sini
		Buku{
			Nomor:     123456,
			Nama:      "Belajar Koding Lima menit",
			Pengarang: "Achmad Fadli Iskandar",
		},
		Buku{
			Nomor:     789012,
			Nama:      "Belajar Koding Tujuh menit",
			Pengarang: "Adrian Rasyad",
		},
		Buku{
			Nomor:     111111,
			Nama:      "Belajar Koding Sepuluh menit",
			Pengarang: "Turah Bagus",
		},
		Buku{
			Nomor:     222222,
			Nama:      "Belajar Koding Lima Belas menit",
			Pengarang: "Edward enstein",
		},
		Buku{
			Nomor:     345612,
			Nama:      "Belajar Itu Tidak ada yang instan",
			Pengarang: "Orang",
		},
	}
	return bk
}
func GetBuku(w http.ResponseWriter,
	r *http.Request) {
	if r.Method == "GET" {
		bk := NewBuku()
		databuku, err := json.Marshal(bk)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(databuku)
		return
	}
	http.Error(w, "hayo mau ngapain",
		http.StatusNotFound)
}

func main() {
	http.HandleFunc("/buku", GetBuku)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
