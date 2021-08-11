package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	p := Person{"Alice", 22}

	jsonData, _ := json.Marshal(p)
	fmt.Println(string(jsonData))
}
