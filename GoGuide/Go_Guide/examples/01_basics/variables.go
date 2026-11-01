// variables.go - Dichiarazione di variabili e valori zero (slide 17).
package main

import "fmt"

// Variabili a livello di package: qui ":=" NON è ammesso, serve "var".
// Il tipo segue il nome; tutte e tre ricevono il valore zero (false).
var c, python, java bool

func main() {
	// Variabile locale senza inizializzatore: valore zero (0 per gli int).
	var i int
	fmt.Println(i, c, python, java)

	// Con inizializzatori (uno per variabile) il tipo può essere omesso
	// e viene dedotto dal valore a destra.
	var a, b = 1, "due"
	fmt.Println(a, b)

	// Dichiarazione breve ":=" - solo dentro le funzioni.
	count := 10
	fmt.Println(count)

	// Valori zero dei tipi base: nessuna variabile Go è "non inizializzata".
	var (
		n int
		f float64
		s string
		t bool
		p *int
	)
	fmt.Printf("int=%d float64=%g string=%q bool=%t pointer=%v\n", n, f, s, t, p)
}
