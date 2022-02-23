package main

// Back and forth communication between two go routines.
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	go pingPong("GoRoutine 1", ch)
	go pingPong("GoRoutine 2", ch)

	ch <- 1

	fmt.Println("Waiting for the Go Routines to finish")
	wg.Wait()
}

func pingPong(label string, ch chan int) {
	wg.Add(1)
	defer wg.Done()
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel already closed")
			fmt.Printf("Exiting on %s\n", label)
			return
		}

		fmt.Printf("The value received by %s is %d\n", label, val)

		if val == 10 {
			close(ch)
			fmt.Printf("Channel closed by %s\n", label)
			return
		}
		val++
		ch <- val
	}
}
