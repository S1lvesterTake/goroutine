package main

import (
	"fmt"
	"time"
)

func main() {

	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		message1 <- "1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		message2 <- "2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-message1:
			fmt.Println("channel 1 get message", msg1)
		case msg2 := <-message2:
			fmt.Println("channel 2 get message", msg2)
		}
	}
}
