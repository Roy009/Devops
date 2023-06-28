package main

import (
	"fmt"
	"sort"
	"strings"
)

// Standard libraries of go
func main(){
	greeting := "hello there friends!"
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "hello"))
	fmt.Println(strings.Split(greeting, " "))

	ages := []int {43,67,89,45,62,12}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 67)
	fmt.Println(index)

	name := []string{"Rohit","Aryan","Roy","David","Sam"}
	sort.Strings(name)
	fmt.Println(name)
	fmt.Println(sort.SearchStrings(name, "Roy"))

}