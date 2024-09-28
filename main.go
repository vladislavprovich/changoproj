package main

import (
	"fmt"
	"sync"
	"time"
)

func prod(id int32) {
	for i := int32(1); i <= 5; i++ {
		id = i
		fmt.Println("Prod start job", id)
		time.Sleep(2 * time.Second)
		fmt.Println("Prod end job", id)
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		prod(1)
	}()

	wg.Wait()
	fmt.Println("Prod end all job")
}
