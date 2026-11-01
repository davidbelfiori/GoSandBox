// arrays.go - Array a dimensione fissa (slide 28).
package main

import "fmt"

func main() {
	// La lunghezza fa parte del tipo: [2]string e [3]string sono tipi
	// diversi.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes, len(primes))

	// Un array è un valore: l'assegnazione ne copia tutti gli elementi
	// (diverso da C, dove un array decade in puntatore).
	b := primes
	b[0] = 99
	fmt.Println(primes[0], b[0]) // 2 99
}
