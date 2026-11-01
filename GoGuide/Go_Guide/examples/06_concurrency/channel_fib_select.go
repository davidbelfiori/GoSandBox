// channel_fib_select.go - Fibonacci con select e channel quit (slide 57-58).
package main

import "fmt"

func fibonacci(ch, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case ch <- x: // invia il numero di Fibonacci corrente
			x, y = y, x+y
		case <-quit: // riceve il segnale di terminazione
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	ch := make(chan int) // channel senza buffer
	quit := make(chan int)

	go func() { // goroutine anonima: il consumatore
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch) // riceve un valore di Fibonacci
		}
		quit <- 0 // dice a fibonacci di fermarsi
	}()

	fibonacci(ch, quit) // eseguita nella goroutine principale
}
