// channel_range.go - range su un channel (slide 55).
package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		// Senza close, il range sotto non terminerebbe mai e main
		// andrebbe in deadlock dopo aver letto i cinque valori.
		close(ch)
	}()

	// Riceve ripetutamente finché il channel è chiuso e svuotato.
	for elem := range ch {
		fmt.Println(elem)
	}
	fmt.Println("channel chiuso, ciclo terminato")
}
