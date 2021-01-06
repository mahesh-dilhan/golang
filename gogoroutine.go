package main

import ("fmt"
        "time"
       )


func main() {
 routine("direct",3)
 go routine("goroutine",3)

//time.Sleep(time.Second)
go func(msg string) {
  fmt.Println(msg)
}("running")

time.Sleep(time.Second)
} 

func routine(from string, i int) {
  for k := 0; k <= i; k++ {
   fmt.Println(from, ":", k)
 }
}
