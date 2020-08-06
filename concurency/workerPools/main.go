package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	const cpuNumber = 1
	const fibNumber = 45
	start := time.Now()

	jobs := make(chan int, 100)
	results := make(chan int, 100)


	for i := 0; i < cpuNumber; i++ {
		go worker(jobs, results)
	}


	for i := 0; i < fibNumber; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < fibNumber; i++ {
		fmt.Println(<-results)
	}

	elapsed := time.Since(start)
	log.Printf("fibNumber: %d took %s [CPU: %d]", fibNumber, elapsed, cpuNumber)

}

func worker(jobs <-chan int, results chan<- int) {
	//work until there are jobs
	for  range jobs {
		results <- fib(40)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
