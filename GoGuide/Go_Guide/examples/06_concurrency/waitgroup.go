// waitgroup.go - sync.WaitGroup per attendere un gruppo di goroutine
// (slide 64).
package main

import (
	"fmt"
	"sync"
)

func worker(id int, results []int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Ogni goroutine scrive in una posizione DIVERSA della slice: non c'è
	// condivisione dello stesso elemento, quindi non serve un lock.
	results[id] = id * id
}

func main() {
	results := make([]int, 5)
	var wg sync.WaitGroup

	for i := range results {
		wg.Add(1) // un'unità in più da attendere
		go worker(i, results, &wg)
	}
	wg.Wait() // attende che tutti abbiano chiamato Done

	// Dopo Wait tutte le scritture sono visibili a main.
	fmt.Println(results) // [0 1 4 9 16]
}
