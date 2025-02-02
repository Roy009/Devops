package main

import "fmt"

func worker(ch chan string) {
	ch <- "Hello from worker!"
}

func main() {
	ch := make(chan string)

	go worker(ch)

	message := <-ch
	fmt.Println("Message Received: ", message)
}
