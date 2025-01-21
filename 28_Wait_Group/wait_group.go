package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitgroup sync.WaitGroup

	waitgroup.Add(2)

	go func() {
		write("Hello world from main")
		waitgroup.Done() // Remove one pilled item on the waitgroup

	}()

	go func() {
		write("Go programming course")
		waitgroup.Done() // Remove one pilled item on the waitgroup
	}()

	waitgroup.Wait()
}

func write(text string) {
	for i := 0; i < 6; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
