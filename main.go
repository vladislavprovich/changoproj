package main

import (
	"fmt"
	"time"
)

//func prod(jobs chan<- int32, count int32) {
//	var i int32
//	for i = 1; i <= count; i++ {
//		fmt.Println("Prod sent the work", i)
//		jobs <- i
//		time.Sleep(100 * time.Millisecond)
//	}
//	close(jobs)
//}
//
//func user(jobs <-chan int32, done chan<- bool) {
//	for job := range jobs {
//		fmt.Println("User has done the work", job)
//		time.Sleep(100 * time.Millisecond)
//	}
//	done <- true
//}

func main() {
	jobs := make(chan int32, 5)    // Buffered channel
	workers := make(chan int32, 5) // Buffered channel
	done := make(chan bool)

	go func() {
		for w := int32(1); w <= 5; w++ {
			jobs <- w
			fmt.Println("Надіслали роботу", w)
			time.Sleep(1 * time.Second)
		}
		close(jobs)
	}()

	go func() {
		for a := range jobs {
			workers <- a
			fmt.Println("Виконали роботу", a)
			time.Sleep(2 * time.Second)
		}
		close(workers)
		done <- true
	}()

	//go prod(jobs, 5)
	//go user(jobs, done)

	for {
		select {
		case j, ok := <-jobs:
			if ok {
				fmt.Println("Received job from producer:", j)
			} else {
				fmt.Println("channel closed")
			}
		case job, ok := <-workers:
			if ok {
				fmt.Println("User has completed job:", job)
			} else {
				fmt.Println("channel closed")
			}
		case <-done:
			{
				fmt.Println("All jobs are done.")
				return
			}

		}
		if workers == nil && done == nil {
			break
		}
	}

}

//new brach
