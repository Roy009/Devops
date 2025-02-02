package main

import (
	"fmt"
	"maps" 
)
func main(){

	// strings
	var name1 string = "Web";
	var name2 = "React";
	var name3 string;
	fmt.Println(name1, name2, name3);

	// change value
	name1 = "app";
	name3 = "android";
	fmt.Println(name1, name2, name3);
	
	// without var and datatype
	name4 := "JS";
	fmt.Println(name4)

	// Integers
	var num1 int = 23;
	var num2 = 34;
	num3 := 45
	fmt.Println(num1, num2, num3);

	// bits and memory
	var num4 int8 = 23
	fmt.Println(num4)

	// non negative values
	var num5 uint = 27
	fmt.Println(num5)

	// floats
	var fnum1 float32 = 23.54;
	fnum2 := 34.54;
	fmt.Println(fnum1, fnum2)

	// formated string
	fmt.Printf("My name is %v and my age is %v \n", name1, num1);
	fmt.Printf("My name is %q and My age is %v\n", name1, num1);
	fmt.Printf("Type of name variable is %T\n", name1)
	fmt.Printf("Float formatted %f\n", fnum1)
	fmt.Printf("Float formatted %0.2f\n", fnum2)
	// Sprintf()
	var str = fmt.Sprintf("My name is %v and my age is %v \n", name1, num1)
	fmt.Println("The saved string is: ", str)
}
