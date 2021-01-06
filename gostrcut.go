package main

import "fmt"

func main(){

 fmt.Println(person{"Mahesh",38})
 fmt.Println(person{name: "Mahesh", age: 38})
 fmt.Println(person{name: "Mahesh"})
 //function
 fmt.Println(newPerson("Mahesh"))
}

type person struct {
  name string
  age  int

}

func newPerson(name string) *person {

 p := person{name:name}
 p.age = 38

 return &p

}

