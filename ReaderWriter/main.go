package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock    sync.RWMutex
	balance int
	wg      sync.WaitGroup
)

func main() {
	balance = 2
	fmt.Println("Reader Writer problem")
	wg.Add(5)
	go read()
	go read()
	go read()
	go writer()
	go read()
	wg.Wait()
}

func read() {
	lock.RLock()
	fmt.Println("Read Lock acquired")
	//fmt.Printf("The value of balance is %d\n", balance)
	time.Sleep(time.Second * 3)
	lock.RUnlock()
	fmt.Println("Read Lock released")
	wg.Done()
}

func writer() {
	lock.Lock()
	fmt.Println("Write Lock acquired")
	balance = balance * 2
	time.Sleep(time.Second * 1)
	lock.Unlock()
	fmt.Println("Write Lock released")
	wg.Done()
}
