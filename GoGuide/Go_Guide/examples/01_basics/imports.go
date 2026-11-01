// imports.go - Import "fattorizzato" tra parentesi (slide 12) e visibilità
// dei nomi (slide 11).
package main

import (
	"fmt"
	"math"
)

func main() {
	// math.Sqrt restituisce un float64; %g sceglie la forma compatta
	// più breve tra %e (esponenziale) e %f (decimale).
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// Un nome esportato (iniziale maiuscola) è accessibile dall'esterno:
	// math.Pi è visibile, mentre math.pi non esiste e l'uso di math.pi
	// produce l'errore di compilazione "undefined: math.pi".
	fmt.Println("math.Pi =", math.Pi)
}
