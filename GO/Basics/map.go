package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string] string)
	m["name"] = "Roy"
	m["age"] = "25"
	fmt.Println(m["name"], m["age"])

	delete(m, "name")
	clear(m)
	fmt.Println(m["name"], m["age"])
	details := map[string] int {"id" : 1, "age" : 24 }
	details1 := map[string] int {"id" : 1, "age" : 24 }
	details2 := map[string] int {"id" : 1, "age" : 24 }
	fmt.Println(details)
	fmt.Println(maps.Equal(details1, details2))

	k, ok := details["id"]

	fmt.Println(k)
	if ok {
		fmt.Println(k)
	} else {
		fmt.Println("Key not found")	
	}
	
}