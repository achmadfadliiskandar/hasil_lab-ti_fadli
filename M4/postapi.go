package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Mahasiswa struct {
	ID   int    `json:"id"`
	NPM  int    `json:"npm"`
	Name string `json:"name"`
}

// post mahasiswa
func PostMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Mhs Mahasiswa

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Mhs); err != nil {
				log.Fatal(err)
			}
		} else {
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			getNpm := r.PostFormValue("npm")
			npm, _ := strconv.Atoi(getNpm)
			name := r.PostFormValue("name")
			Mhs = Mahasiswa{
				ID:   id,
				NPM:  npm,
				Name: name,
			}
		}
		dataMahasiswa, _ := json.Marshal(Mhs)
		w.Write(dataMahasiswa)
		return
	}
	http.Error(w, "hayo mau ngapain",
		http.StatusNotFound)
	return
}

func main() {
	http.HandleFunc("/mahasiswa", PostMahasiswa)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
