package main

import "fmt"

func main() {
	// Array Declaration
	var ages [3]int = [3]int{20, 25, 30}
	ages[1] = 34
	fmt.Println(ages)
	fmt.Println(len(ages))
	name := [4]string{"Rohit", "Roy", "David", "Ronaldo"}
	fmt.Println(name, len(name))

	// slices
	var score = []int{100, 40, 50}
	score[1] = 0
	score = append(score, 80)
	fmt.Println(score)

	// slice Ranges
	rangeOne := name[1:3]
	rangeTwo := name[2:]
	rangeThree := name[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Vini Jr")
	fmt.Println(rangeOne)

	// Slice array
	var num = make([]int, 3, 100)
	num[0] = 10
	num[1] = 15
	num[2] = 18
	num = append(num, 20)

	nums := make([]int, 3)
	copy(nums, num)
	fmt.Println(nums)
	fmt.Println(num, len(num), cap(num))
}
