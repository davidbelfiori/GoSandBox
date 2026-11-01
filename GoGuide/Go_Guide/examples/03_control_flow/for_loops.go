// for_loops.go - L'unico costrutto di ciclo di Go: for (slide 20-21).
package main

import "fmt"

func main() {
	// Forma completa: init; condizione; post. Niente parentesi attorno
	// ai tre componenti, graffe sempre obbligatorie.
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("somma 0..9 =", sum)

	// Senza init e post il for si comporta come un "while".
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("primo valore >= 1000:", sum)

	// Senza condizione il ciclo è infinito; si esce con break (o return).
	n := 0
	for {
		n++
		if n == 5 {
			break
		}
	}
	fmt.Println("uscito dal ciclo infinito con n =", n)
}
