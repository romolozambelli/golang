package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 2)

	channel <- "Hello World Main"
	channel <- "Golan Programming"
	// channel <- "Golan Deadlock"

	mensagem1 := <-channel
	mensagem2 := <-channel

	fmt.Println(mensagem1)
	fmt.Println(mensagem2)
}

func write(text string, canal chan string) {
	for i := 0; i < 6; i++ {

		canal <- text

		time.Sleep(time.Second)
	}

	close(canal)
}
