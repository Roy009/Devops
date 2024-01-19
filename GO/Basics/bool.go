package main

import(
	"fmt"
)

func main(){
	var age int = 13
	var name string = "Marketing user"

	// fmt.Println(age <= 30)
	// fmt.Println(age >= 20)
	// fmt.Println(age == 25)
	// fmt.Println(age != 24)
	if age > 18 {
		fmt.Println("You can Drive")
	}else if age > 20 {
		fmt.Println("You cannot Drive")
	}else {
		fmt.Printf("Your age is %v You cannot drive", age )
	}

	if (name == "Sales person")|| (name == "Marketing User") {
		fmt.Println("This is a Bussiness profile")
	}else if (name ==  "SDE") || (name == "SRE") {
		fmt.Println("This is a Technical Profile")
	}else{
		fmt.Println("This is a general profile")
	}
	
}

