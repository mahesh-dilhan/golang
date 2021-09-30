package main

import (
	"fmt"
)

type FullNameType func(fname string, lname string) string

type PPerson struct {
	fullnsmetye FullNameType
	firstname   string
	lastname    string
}

func main() {

	p := PPerson{
		firstname: "Mahesh",
		lastname:  "Wijekoon",
		fullnsmetye: func(fname string, lname string) string {
			return fname + " " + lname
		},
	}
	fmt.Println(p.fullnsmetye(p.firstname, p.lastname))
}
