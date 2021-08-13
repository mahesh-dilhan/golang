package main

import "fmt"

type Vaccine struct {
	vname string
}

type CovidCountry struct {
	Vaccine
	postiveCases int
	name         string
}

func main() {
	c := &CovidCountry{
		name:         "USA",
		postiveCases: 100,
	}
	fmt.Println(*c)
	c.vname = "Mordana"
	fmt.Println(*c)
}
