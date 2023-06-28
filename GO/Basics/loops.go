package main
import (
	"fmt"
)

func main(){
	
	x := 0
	for x < 5 {
		fmt.Println("value of x is: ", x)
		x++
	}

	name := []string{"Rohit","Aryan","Roy","David"}

	for i:= 0; i < len(name); i++{
		fmt.Println("The name is ", name[i])
	}

	for index, value := range name {
		fmt.Printf("The position at index %v is %v\n", index, value)
	}

}