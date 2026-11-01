// anonymous.go - Funzioni anonime (slide 37).
package main

import "fmt"

func main() {
	// Funzione anonima assegnata a una variabile.
	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(3, 4)) // 7

	// Funzione anonima definita e invocata subito: le parentesi finali
	// () eseguono la chiamata.
	func(msg string) {
		fmt.Println("invocata subito:", msg)
	}("ciao")
}
