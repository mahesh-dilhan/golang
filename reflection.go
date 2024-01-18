// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

type engine interface {
	milesleft() int
}
type gasEngine struct {
	mpg     int
	gallons int
}

type electricEngine struct {
	mpkwh int
	kwh   int
}

func (e gasEngine) milesleft() int {
	return e.mpg * e.gallons
}

func (e electricEngine) milesleft() int {
	return e.mpkwh * e.kwh
}

func whatEngine(t interface{}) string {
	switch t.(type) {
	case gasEngine:
		return "gasEngine"
	default:
		return "electricEngine"
	}
}

func canweMakeIt(e engine, miles int) string {
	var state = (miles <= e.milesleft())
	if !state {
		fmt.Println("We need a", whatEngine(e))
	}

	if state {
		return "We can make it..!"
	}

	return "Opps, we need to find solution"
}

func main() {
	var useengine1 gasEngine = gasEngine{2, 10}
	fmt.Println(canweMakeIt(useengine1, 50))

	var useengine2 electricEngine = electricEngine{20, 10}
	fmt.Println(canweMakeIt(useengine2, 50))
}
