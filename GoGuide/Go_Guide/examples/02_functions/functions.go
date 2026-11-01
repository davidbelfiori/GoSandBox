// functions.go - Funzioni: parametri, tipo dopo il nome, forma abbreviata
// (slide 13).
package main

import "fmt"

// Il tipo segue il nome del parametro; anche il tipo di ritorno è in coda.
func add(x int, y int) int {
	return x + y
}

// Forma abbreviata: parametri consecutivi dello stesso tipo.
func addShort(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(addShort(10, 20))
}
