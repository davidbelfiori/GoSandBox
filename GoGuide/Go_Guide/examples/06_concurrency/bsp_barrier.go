// bsp_barrier.go - Bulk Synchronous Parallel con WaitGroup come barriera
// (slide 64).
//
// Pattern BSP: le goroutine lavorano indipendentemente (fase di calcolo),
// poi attendono tutte a una barriera; solo dopo si passa alla fase
// successiva, in cui i dati prodotti sono visibili a tutti.
package main

import (
	"fmt"
	"sync"
)

func main() {
	const n = 4
	partial := make([]int, n)

	// Fase 1: ogni goroutine calcola un risultato parziale.
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			partial[i] = (i + 1) * 10
		}()
	}
	wg.Wait() // BARRIERA

	// Scambio dati alla barriera: ogni goroutine della fase 2 legge
	// i risultati parziali di tutte le altre.
	total := 0
	for _, p := range partial {
		total += p
	}
	fmt.Println("totale dopo la fase 1:", total) // 100

	// Fase 2: ogni goroutine usa il totale globale.
	shares := make([]float64, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			shares[i] = float64(partial[i]) / float64(total)
		}()
	}
	wg.Wait() // seconda barriera
	fmt.Println("quote:", shares)
}
