// mutex.go - sync.Mutex per la mutua esclusione (slide 63).
package main

import (
	"fmt"
	"sync"
)

// Counter protegge il campo n con un mutex.
type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	// defer garantisce lo Unlock anche in caso di return anticipato o panic.
	defer c.mu.Unlock()
	c.n++
}

func main() {
	var c Counter
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				c.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Println("counter =", c.n) // sempre 100000
}
