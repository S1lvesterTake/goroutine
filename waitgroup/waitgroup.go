package main

import (
	"fmt"
	"sync"
	"time"
)

func sender(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Sender %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Sender %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go sender(i, &wg)
	}

	wg.Wait()
	fmt.Println("All goroutine done")
}
