// closure_fib.go - Fibonacci con una closure (slide 39).
package main

import "fmt"

// fibonacci restituisce una funzione che, a ogni chiamata, restituisce
// il numero di Fibonacci successivo. Lo stato (x, y) vive nella closure.
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		// Assegnazione parallela: il lato destro è valutato per intero
		// prima di assegnare, quindi non serve una variabile temporanea.
		x, y = y, x+y
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
