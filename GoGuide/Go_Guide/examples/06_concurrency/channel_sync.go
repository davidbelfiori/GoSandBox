// channel_sync.go - Un channel come meccanismo di sincronizzazione
// (slide 50).
//
// Esempio aggiuntivo: main attende la fine di una goroutine ricevendo da un
// channel "done". Non servono lock né condition variable.
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("worker: inizio lavoro...")
	time.Sleep(100 * time.Millisecond) // lavoro simulato
	fmt.Println("worker: fatto")
	done <- true // segnala il completamento
}

func main() {
	done := make(chan bool)
	go worker(done)

	// La ricezione blocca main finché il worker non invia.
	<-done
	fmt.Println("main: il worker ha terminato")
}
