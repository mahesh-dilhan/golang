package main

import (
	"fmt"
)

type FullNameType func(fname string, lname string) string

type PPerson struct {
	firstname   string
	lastname    string
	fullnsmetye FullNameType
}

func main() {

	p := PPerson{
		firstname: "Mahesh",
		lastname:  "Wijekoon",
		fullnsmetye: func(fname string, lname string) string {
			return fname + lname
		},
	}
	fmt.Println(p)
}
