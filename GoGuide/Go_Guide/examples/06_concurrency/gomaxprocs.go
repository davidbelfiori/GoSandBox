// gomaxprocs.go - Goroutine leggere e thread del sistema operativo
// (slide 47-48).
//
// runtime.GOMAXPROCS(0) restituisce il numero massimo di thread OS che
// possono eseguire codice Go contemporaneamente, senza modificarlo. Il
// valore dipende dalla macchina, quindi viene stampato solo un controllo
// di coerenza per mantenere l'output identico ovunque.
//
// Qui si avviano 10000 goroutine, molte più dei thread disponibili: lo
// scheduler di Go le multiplexa sui thread OS. Con 10000 thread OS reali
// il costo in memoria e tempo di creazione sarebbe molto maggiore.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("GOMAXPROCS(0) >= 1:", runtime.GOMAXPROCS(0) >= 1)

	const n = 10000
	finished := make([]bool, n) // una posizione per goroutine: nessuna race
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			finished[i] = true
		}()
	}
	wg.Wait()

	count := 0
	for _, f := range finished {
		if f {
			count++
		}
	}
	fmt.Println("goroutine completate:", count)
}
