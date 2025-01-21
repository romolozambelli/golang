package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplex(write("Hello World"), write("GO land Multiplex"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplex(channelInput1, channelInput2 <-chan string) <-chan string {
	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-channelInput1:
				outputChannel <- message

			case message := <-channelInput2:
				outputChannel <- message
			}
		}

	}()
	return outputChannel

}

func write(text string) <-chan string {

	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received value = %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return channel
}
