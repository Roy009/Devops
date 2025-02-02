package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i:= 1; i <=5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Start a go routine
	go printNumbers()

	// Main function comntinues to run
	fmt.Println("Main function running...")

	// Give goroutine some time to run
	time.Sleep(3 * time.Second)
}