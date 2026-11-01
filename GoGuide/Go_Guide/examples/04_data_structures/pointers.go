// pointers.go - Puntatori (slide 27).
package main

import "fmt"

// Go è call-by-value: la funzione riceve una copia del puntatore, che però
// punta alla stessa variabile del chiamante.
func incrementByValue(n int) { n++ }
func incrementByPointer(n *int) {
	*n++ // dereferenziazione: modifica la variabile originale
}

// È sicuro restituire l'indirizzo di una variabile locale: il runtime la
// alloca sullo heap (escape) e il garbage collector la libera quando
// nessuno la referenzia più. In C sarebbe un dangling pointer.
func newCounter() *int {
	c := 100
	return &c
}

func main() {
	var p *int
	fmt.Println("valore zero del puntatore:", p) // <nil>

	i := 1
	p = &i          // p punta a i
	fmt.Println(*p) // 1
	*p = 2          // equivale a i = 2
	fmt.Println(i)  // 2

	incrementByValue(i)
	fmt.Println("dopo incrementByValue:", i) // invariato
	incrementByPointer(&i)
	fmt.Println("dopo incrementByPointer:", i)

	q := newCounter()
	*q++
	fmt.Println("contatore sullo heap:", *q)

	// Nessuna aritmetica dei puntatori: p++ o p+1 non compilano.
}
