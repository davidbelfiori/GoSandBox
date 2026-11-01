// goroutine_termination.go - ERRORE COMUNE: main termina prima delle
// goroutine (slide 47).
//
// Quando main ritorna, il programma termina immediatamente senza attendere
// le altre goroutine. L'output di questo programma è NON deterministico:
// dopo la riga di main possono comparire da 0 a 3 righe "goroutine N"
// (anche prima della riga di main, in ordine qualsiasi). Nelle prove fatte
// sono comparse quantità diverse da un'esecuzione all'altra: non c'è
// alcuna garanzia.
package main

import "fmt"

func say(id int) {
	fmt.Println("goroutine", id)
}

func main() {
	for i := 1; i <= 3; i++ {
		go say(i) // avvia la goroutine e prosegue subito
	}
	fmt.Println("main termina senza attendere")
	// Nessuna sincronizzazione: il programma esce qui.
	// Soluzione: vedi goroutine_wait.go (sync.WaitGroup).
}
