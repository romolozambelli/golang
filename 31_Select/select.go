package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Channel 2"
		}
	}()

	for {
		select {
		case message1 := <-channel1:
			fmt.Println(message1)
		case message2 := <-channel2:
			fmt.Println(message2)
		}
	}

	// channel <- "Hello World Main"
	// channel <- "Golan Programming"
	// // channel <- "Golan Deadlock"

	// mensagem1 := <-channel
	// mensagem2 := <-channel

	// fmt.Println(mensagem1)
	// fmt.Println(mensagem2)
}

func write(text string, canal chan string) {
	for i := 0; i < 6; i++ {

		canal <- text

		time.Sleep(time.Second)
	}

	close(canal)
}
