package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Mahasiswa struct {
	ID   int    `json:"id"`
	NPM  int    `json:"npm"`
	Name string `json:"name"`
}

func NewMahasiswa() []Mahasiswa {
	mhs := []Mahasiswa{
		Mahasiswa{
			ID:   1,
			NPM:  50421055,
			Name: "Adrian",
		},
		Mahasiswa{
			ID:   2,
			NPM:  51419643,
			Name: "Dariwan",
		},
		Mahasiswa{
			ID:   3,
			NPM:  50419832,
			Name: "Ardava",
		},
	}
	return mhs
}
func GetMahasiswa(w http.ResponseWriter,
	r *http.Request) {
	if r.Method == "GET" {
		mhs := NewMahasiswa()
		datamahasiswa, err := json.Marshal(mhs)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(datamahasiswa)
		return
	}
	http.Error(w, "hayo mau ngapain",
		http.StatusNotFound)
}

func main() {
	http.HandleFunc("/mahasiswa", GetMahasiswa)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
