package main

import "fmt"

type Person struct{
	fname, lname string
}

var(
	p1 = Person{"Rohit","Roy"}
	p2 = Person{"Rohit",""}
	p3 = Person{lname: "Roy"}
)

func main(){
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}