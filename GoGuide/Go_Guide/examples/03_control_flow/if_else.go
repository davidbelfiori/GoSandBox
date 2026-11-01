// if_else.go - if con istruzione di inizializzazione (slide 23-24).
package main

import (
	"fmt"
	"math"
)

// pow restituisce x^n se è inferiore a limit, altrimenti stampa un
// messaggio e restituisce limit.
func pow(x, n, limit float64) float64 {
	// v è dichiarata nell'istruzione iniziale dell'if: è visibile nel blocco
	// if e nel relativo else, ma non dopo.
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		// Nota: "} else {" deve stare sulla stessa riga.
		fmt.Printf("%g >= %g\n", v, limit)
	}
	// Qui "v" non esiste più (errore di compilazione se la si usasse).
	return limit
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	// Catena if / else if / else, ciascuno con la propria istruzione
	// opzionale.
	x := 7
	if r := x % 2; r == 0 {
		fmt.Println("pari")
	} else if r == 1 {
		fmt.Println("dispari")
	} else {
		fmt.Println("negativo e dispari")
	}
}
