package main

import (
	"fmt"
	"time"
)

func du(pinger <-chan int, ponger chan<- int) {

	<-pinger
	ponger <- 1
}

func po(pinger chan<- int, ponger <-chan int) {
	<-ponger
	pinger <- 1
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

	go pinger(ping, pong)
	go ponger(ping, pong)

	ping <- 1

	for {
		time.Sleep(time.Second)
	}
}
