package main

import "fmt"

func main(){
	nums := []int {1,2,3,4,5}
	for i, num := range nums {
		fmt.Println("Index --> ", i, " value is --> ", num)
	}
	
	m := map[string] string { "First Name" : "Rohit", "Last Name" : "Roy", "Age" : "24" }

	for k, v := range m{
		fmt.Println("Key --> ", k, " || Value --> ", v)
	}

	for i, s := range "golang" {
		fmt.Println(i, s, i, string(s))
	}
}