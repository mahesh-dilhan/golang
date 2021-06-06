package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 1_000_000; i++ {
		go func(msg string) {
			fmt.Print(msg)
		}(".")
	}

	//time.Sleep(time.Minute)
}
