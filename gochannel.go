package main

import (
	"fmt"
	"time"
)

func du(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping 2")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

func po(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong 2")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <- 1

	}

}

func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {

	ping := make(chan int)
	pong := make(chan int)

	ping2 := make(chan int)
	pong2 := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	go du(ping2, pong2)
	go po(ping2, pong2)

	ping <- 1
	ping2 <- 1

	for {
		time.Sleep(time.Second)
	}
}
