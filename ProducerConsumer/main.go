package main

import (
	"fmt"
	"sync"
)

const producerCount int = 4

var wg sync.WaitGroup
var messages = [][]string{
	{
		"a", "b", "c", "d",
	},
	{
		"e", "f", "g", "h",
	},
	{
		"i", "j", "k", "l",
	},
	{
		"m", "n", "o", "p",
	},
}

func produce(ch chan<- string, index int) {
	defer wg.Done()

	for _, a := range messages[index] {
		ch <- a
	}
}

func consume(ch <-chan string, done chan<- bool) {
	for a := range ch {
		fmt.Printf("Consumed: %s\n", a)
	}
	done <- true
}
func main() {
	ch := make(chan string, 4)
	done := make(chan bool)
	for i := 0; i < producerCount; i++ {
		wg.Add(1)
		go produce(ch, i)
	}
	go consume(ch, done)
	wg.Wait()
	close(ch)
	<-done
	fmt.Println("Completed the process")
}
