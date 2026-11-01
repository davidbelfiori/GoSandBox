// errors.go - Gestione degli errori e interfaccia error (slide 69).
package main

import (
	"errors"
	"fmt"
)

// L'interfaccia predefinita error è:
//
//	type error interface { Error() string }
//
// Qualunque tipo con un metodo Error() string la soddisfa.

// divide segnala la condizione anomala con un errore invece che con
// un'eccezione. L'errore è l'ultimo valore restituito.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// errors.New costruisce un errore base con un messaggio.
		return 0, errors.New("divisione per zero")
	}
	return a / b, nil // nil = nessun errore
}

// Un tipo errore personalizzato: porta con sé dati strutturati.
type rangeError struct {
	value, max int
}

func (e rangeError) Error() string {
	return fmt.Sprintf("valore %d oltre il massimo %d", e.value, e.max)
}

func check(v int) error {
	if v > 10 {
		return rangeError{v, 10}
	}
	return nil
}

func main() {
	result, err := divide(10, 4)
	if err != nil {
		fmt.Println("errore:", err)
		return
	}
	fmt.Println("risultato:", result)

	if _, err := divide(1, 0); err != nil {
		fmt.Println("errore:", err)
	}

	for _, v := range []int{5, 42} {
		if err := check(v); err != nil {
			fmt.Println("check fallito:", err)
		} else {
			fmt.Println("check ok:", v)
		}
	}
}
