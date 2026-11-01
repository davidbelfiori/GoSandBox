// channel_fib.go - Generatore di Fibonacci con goroutine e channel
// (slide 52).
package main

import "fmt"

// fib è un produttore: invia numeri di Fibonacci sul channel c, per
// sempre. Ogni invio blocca finché il main non riceve (channel senza
// buffer), quindi il produttore non va "avanti" rispetto al consumatore.
func fib(c chan int) {
	x, y := 1, 1
	for {
		c <- x
		x, y = y, x+y
	}
}

func main() {
	c := make(chan int)
	go fib(c)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	// Quando main termina, la goroutine fib (bloccata su un invio)
	// viene semplicemente abbandonata: si veda select per una chiusura
	// ordinata.
}
