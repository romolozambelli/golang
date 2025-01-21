package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("Hello World", channel)

	// for {
	// 	message, open := <-channel
	// 	if !open {
	// 		break
	// 	}

	// 	fmt.Println(message)

	// }

	for message := range channel {
		fmt.Println(message)
	}
}

func write(text string, canal chan string) {
	for i := 0; i < 6; i++ {

		canal <- text

		time.Sleep(time.Second)
	}

	close(canal)
}
