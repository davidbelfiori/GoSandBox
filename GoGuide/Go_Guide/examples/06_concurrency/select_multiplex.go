// select_multiplex.go - select su più channel e scelta casuale (slide 56).
package main

import "fmt"

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch1 <- "da ch1"
	ch2 <- "da ch2"

	// Entrambi i case sono pronti: select ne sceglie UNO a caso. L'ordine
	// delle due righe stampate può quindi variare da esecuzione a
	// esecuzione, ma entrambi i messaggi vengono sempre ricevuti.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("ricevuto", msg1)
		case msg2 := <-ch2:
			fmt.Println("ricevuto", msg2)
		}
	}
}
