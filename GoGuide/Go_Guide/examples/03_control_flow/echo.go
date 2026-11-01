// Echo prints its command-line arguments. (slide 22)
//
// Esecuzione: go run echo.go Hello Go World
package main

import (
	"fmt"
	"os"
)

func main() {
	// s e sep sono inizializzate implicitamente alla stringa vuota.
	var s, sep string
	// os.Args è una slice di stringhe; os.Args[0] è il nome del programma,
	// quindi gli argomenti utili partono da 1.
	for i := 1; i < len(os.Args); i++ {
		// Al primo giro sep è "" (nessuno spazio iniziale); poi diventa " ".
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
