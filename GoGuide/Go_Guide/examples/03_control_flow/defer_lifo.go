// defer_lifo.go - Ordine LIFO e valutazione immediata degli argomenti
// (slide 26).
package main

import "fmt"

func countdown() {
	// Le chiamate differite sono impilate: l'ultima registrata è la prima
	// eseguita. Ogni defer cattura il valore corrente di i (argomento
	// valutato subito).
	for i := 1; i <= 3; i++ {
		defer fmt.Println("defer", i)
	}
	fmt.Println("fine del corpo di countdown")
}

func main() {
	x := 1
	// L'argomento x viene valutato qui: la stampa mostrerà 1, non 2.
	defer fmt.Println("x al momento del defer =", x)
	x = 2
	fmt.Println("x ora =", x)

	countdown()
}
