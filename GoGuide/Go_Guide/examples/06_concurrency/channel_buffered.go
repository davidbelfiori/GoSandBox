// channel_buffered.go - Channel con e senza buffer (slide 53).
package main

import "fmt"

func main() {
	// Con buffer di capacità 2: i primi due invii NON bloccano anche se
	// nessuno sta ancora ricevendo.
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("len =", len(ch), "cap =", cap(ch)) // len=2 cap=2

	// Un terzo invio bloccherebbe (buffer pieno) e, non esistendo altre
	// goroutine che ricevano, causerebbe un deadlock. Le ricezioni
	// seguono l'ordine FIFO.
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2

	// Un channel senza buffer (capacità 0) richiede sender e receiver
	// pronti contemporaneamente: l'invio avviene in una goroutine perché
	// altrimenti main si bloccherebbe da solo.
	unbuf := make(chan string)
	go func() { unbuf <- "ping" }()
	fmt.Println(<-unbuf)
}
