// goroutine_wait.go - Versione corretta di goroutine_termination.go:
// main attende le goroutine con sync.WaitGroup (slide 47, 64).
//
// L'ORDINE con cui le goroutine stampano non è garantito: tutte hanno la
// stessa priorità e lo scheduler decide. Quindi l'output varia da
// un'esecuzione all'altra (ma contiene sempre le tre righe).
package main

import (
	"fmt"
	"sync"
)

func say(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrementa il contatore all'uscita, anche in caso di return anticipato
	fmt.Println("goroutine", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // PRIMA di avviare la goroutine, non al suo interno
		go say(i, &wg)
	}
	wg.Wait() // blocca finché il contatore torna a zero
	fmt.Println("main: tutte le goroutine sono terminate")
}
