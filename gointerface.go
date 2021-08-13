package main

import "fmt"

type Vechicle interface {
	forward(x int)
	backword(x int)
}

type Van struct {
	km int
}

func (v *Van) forward(x int) {
	fmt.Println("moving forwar", x)
}

func (v *Van) backword(x int) {
	fmt.Println("moving backword", x)
}

func main() {
	var v Vechicle = &Van{km: 10}
	fmt.Println(v)
}
