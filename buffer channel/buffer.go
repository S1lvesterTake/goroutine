package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 3)

	go func() {
		for {
			i := <-messages
			fmt.Println("menerima data ", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("mengirim data ", i)
		messages <- i
	}
}
