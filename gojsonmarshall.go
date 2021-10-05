package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Bird struct {
	Species     string
	Description string
}

func main() {

	p := Person{"Alice", 22}
	jsonData, _ := json.Marshal(p)
	fmt.Println(string(jsonData))

	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)

}
