// rwmutex.go - sync.RWMutex: molti lettori oppure un solo scrittore
// (slide 63).
package main

import (
	"fmt"
	"sync"
)

// Store è una cache chiave-valore letta spesso e scritta di rado.
type Store struct {
	mu   sync.RWMutex
	data map[string]int
}

// Get usa un lock di LETTURA: più Get possono procedere insieme.
func (s *Store) Get(k string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.data[k]
	return v, ok
}

// Set usa il lock di SCRITTURA: accesso esclusivo (blocca lettori e altri
// scrittori).
func (s *Store) Set(k string, v int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[k] = v
}

func main() {
	s := &Store{data: make(map[string]int)}
	s.Set("a", 1)

	var wg sync.WaitGroup
	// 10 lettori concorrenti.
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if v, ok := s.Get("a"); !ok || v < 1 {
				fmt.Println("valore inatteso")
			}
		}()
	}
	// 1 scrittore concorrente.
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Set("b", 2)
	}()
	wg.Wait()

	v, _ := s.Get("b")
	fmt.Println("b =", v)
}
