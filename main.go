package main

import (
	"fmt"
	"time"
)

func prod(jobs chan<- int32, count int32) {
	var i int32
	for i = 1; i <= count; i++ {
		fmt.Println("Prod sent the work", i)
		jobs <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(jobs)
}

func user(jobs <-chan int32, done chan<- bool) {
	for job := range jobs {
		fmt.Println("User has done the work", job)
		time.Sleep(100 * time.Millisecond)
	}
	done <- true
}

func main() {
	jobs := make(chan int32)
	done := make(chan bool)

	go prod(jobs, 10) // Buffered channel
	go user(jobs, done)

	<-done // wait user done jobs
	fmt.Println("All jobs done")
}
