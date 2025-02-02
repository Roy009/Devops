package main


import "fmt"

func main(){
	i := 24
	p := &i
	fmt.Println("Value of i is ", i)
	fmt.Println("Address of p is ", p)
	fmt.Println("Address of p is ", *p)
}