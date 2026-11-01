// race_counter.go - ERRORE COMUNE: data race su una variabile condivisa
// (slide 47, 63).
//
// counter++ NON è atomico (legge, incrementa, scrive). Con più goroutine
// non sincronizzate alcuni aggiornamenti si perdono. Il risultato è
// NON deterministico e quasi sempre inferiore a 100000. (In rari casi su
// macchine con un solo thread potrebbe risultare esatto.)
package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++ // DATA RACE: nessuna mutua esclusione
			}
		}()
	}
	wg.Wait()
	fmt.Println("atteso 100000, ottenuto", counter)
	// Correzione: vedi mutex.go (sync.Mutex).
}
