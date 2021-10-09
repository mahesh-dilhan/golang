package main

import "fmt"

type Body struct {
	Msg interface{}
}

func main() {
	b := Body{"Hello"}
	fmt.Printf("%#v %T\n", b.Msg, b.Msg)

	b = Body{100}
	fmt.Printf("%#v %T\n", b.Msg, b.Msg)
}
