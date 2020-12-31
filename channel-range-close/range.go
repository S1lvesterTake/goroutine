package main

import "fmt"

func main() {
	channel := make(chan int, 5)
	channel <- 1
	channel <- 2
	channel <- 3
	close(channel)

	for element := range channel {
		fmt.Println(element)
	}

	channel <- 3 //comment this to avoid error
}
