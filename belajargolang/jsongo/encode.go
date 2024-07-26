package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Nama     string `json:"Nama"`
	Semester int
}

func main() {
	var object = []User{{"Achmad Fadli", 2}, {"Semester", 3}}
	var jsonData, err = json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
