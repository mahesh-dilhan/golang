package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "covidservice-api", log.LstdFlags)
	gh := greethander.NewGreet(l)

	sm := http.NewServeMux()
	sm.Handle("/", gh)
	http.ListenAndServe(":9090", sm)
}
