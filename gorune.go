package main

import (
	"fmt"
	"strings"
)

func main() {

	r := []rune("Hello")
	r[0] = 'G'
	buf := string(r)
	fmt.Println(buf)

	s := fmt.Sprintf("%c%c, world!", 72, 'i')
	fmt.Println(s) // "Hi, world!"

	f := func(r rune) rune {
		return r + 1
	}
	fmt.Println(strings.Map(f, "ab"))

	sp := strings.Split("a,b", ",")
	fmt.Printf("%T ", sp)
	fmt.Println("")
	for _, VALUE := range sp {
		fmt.Println(VALUE)
	}
}
