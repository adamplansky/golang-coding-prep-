package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ProcessJob(5)
	fmt.Println(time.Since(start))
}

func ProcessJob(jobs int) {
	sem := make(chan int)
	res := make(chan int)
	for i := 0; i < jobs; i++ {
		go SomeHeavyWorkload(sem, res)
		sem <- i
	}

	for i := 0; i < jobs; i++ {
		goJob := <-res
		fmt.Printf("goJob [%d] complete!\n", goJob)
	}
}

func SomeHeavyWorkload(sem, res chan int) {
	jobNumber := <-sem
	fmt.Printf("processing heaving workload %d\n", jobNumber)
	time.Sleep(time.Millisecond * 300)
	fmt.Printf("Job [%d] complete!\n", jobNumber)
	res <- jobNumber
}

//func ProcessJob(jobs int) {
//	sem := make(chan int)
//	for i := 0; i < jobs; i++ {
//		go SomeHeavyWorkload(i, sem)
//		<-sem
//	}
//}
//
//func SomeHeavyWorkload(jobNumber int, sem chan int) {
//	fmt.Printf("processing heaving workload %d\n", jobNumber)
//	time.Sleep(time.Millisecond * 300)
//	fmt.Printf("Job [%d] complete!\n", jobNumber)
//	sem <- jobNumber
//}
