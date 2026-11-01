// timers.go - Timer e timeout con select (slide 56, 60).
package main

import (
	"fmt"
	"time"
)

func main() {
	// Attesa di una durata: il timer espone un channel C su cui riceve
	// un valore quando scade.
	start := time.Now()
	timer := time.NewTimer(200 * time.Millisecond)
	<-timer.C // blocca fino allo scadere
	fmt.Println("timer 1 scaduto, trascorsi almeno 200ms:",
		time.Since(start) >= 200*time.Millisecond)

	// Un timer si può annullare con Stop() prima che scatti.
	timer2 := time.NewTimer(time.Second)
	stopped := timer2.Stop()
	fmt.Println("timer 2 annullato:", stopped)

	// Timeout: select attende un risultato oppure il timer, a seconda di
	// quale arrivi prima.
	result := make(chan string)
	go func() {
		time.Sleep(500 * time.Millisecond) // operazione lenta
		result <- "risultato"
	}()

	// Anche il timeout è un timer: il suo channel C si "attiva" allo
	// scadere dei 100ms. Vince il case che diventa pronto per primo.
	timeout := time.NewTimer(100 * time.Millisecond)
	select {
	case r := <-result:
		fmt.Println(r)
		timeout.Stop()
	case <-timeout.C:
		fmt.Println("timeout: l'operazione è troppo lenta")
	}
}
