package main

import "fmt"

func main(){

 fmt.Println(person{"Mahesh",38})
 fmt.Println(person{name: "Mahesh", age: 38})
 fmt.Println(person{name: "Mahesh"})
 //function
 fmt.Println(newPerson("Mahesh"))
 r := shaperect{width: 10, height: 10}
  fmt.Println(r.perimeter())
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

type shaperect struct {
 width, height int

}


func (r *shaperect) perimeter() int {
   return 2*r.width + 2*r.height
 }
