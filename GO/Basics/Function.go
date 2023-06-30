package main

import (
	"fmt"
	"math"
)

func greetings(str string){
	fmt.Printf("Welcome to Go %v \n",str)
}

func goodBye(str string){
	fmt.Printf("Good Bye %v \n", str)
}

func addTwoNum(num1 int, num2 int){
	fmt.Printf("The Sum of two numbers are %v", num1 + num2)
}

func cycleNames(n []string, f func(string)){
	for _, v := range n {
		f(v)
	}
}

// func functionName(parameters) return type { Then logic }
func circleArea(r float64) float64 {
	return math.Pi * r * r
}



func main(){
	// greetings("Rohit")
	// goodBye("Rohit")
	// addTwoNum(1267,9089)
	cycleNames([]string{"Hero","Hercules"},greetings)
	cycleNames([]string{"Hero","Hercules"},goodBye)

	num1 := circleArea(2.7)
	num2 := circleArea(3.6)
	fmt.Printf("Area of circle-1 is %0.3f \n", num1)
	fmt.Printf("Area of circle-2 is %0.2f \n", num2)
	addTwoNum(int(num1), int(num2))
}