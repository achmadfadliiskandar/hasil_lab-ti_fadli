package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name" : "Adrian Nova", "Age": 20}`
	var jsonData = []byte(jsonString)
	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user : ", data.FullName)
	fmt.Println("umur :", data.Age)
}
