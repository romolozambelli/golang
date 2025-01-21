package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Hello world from main")
	write("Go programming course")

}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
