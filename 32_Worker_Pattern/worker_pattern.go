package main

import "fmt"

func main() {

	fmt.Println("-- Worker Pattern --")

	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i

	}
	close(tasks)

	for i := 0; i < 45; i++ {
		results := <-results
		fmt.Println(results)
	}
}

func returnNumberFibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return returnNumberFibonacci(position-2) + returnNumberFibonacci(position-1)
}

func worker(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		results <- returnNumberFibonacci(number)
	}

}
