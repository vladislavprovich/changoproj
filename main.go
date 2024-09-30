package main

import (
	"log"
	"sync"
)

type unit struct {
	energy int32
	mu     sync.Mutex
}

func (u *unit) spend(spend int32) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.energy-spend < 0 {
		return nil
	}

	u.energy = u.energy - spend
	log.Printf("Затрачено %d енергії, залишилося %d ", spend, u.energy)
	return nil
}

func jump(u *unit) {
	if err := u.spend(3); err != nil {
		log.Printf("Не вистачає енергії")
	}

}

func sprint(u *unit) {
	if err := u.spend(4); err != nil {
		log.Printf("Не вистачає енергії")
	}

}

func main() {
	var (
		unitsetings = unit{energy: 5}
		wg          sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		jump(&unitsetings)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sprint(&unitsetings)
	}()

	wg.Wait()

}

//task mutex
