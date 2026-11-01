// closure.go - Closure con stato privato (slide 38).
package main

import "fmt"

// intSeq restituisce una funzione. La funzione anonima "cattura" la
// variabile i dell'ambiente che la racchiude: i sopravvive alle
// chiamate di intSeq ed è privata di quella particolare closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	// Una nuova closure ha una propria copia indipendente di i.
	newInts := intSeq()
	fmt.Println(newInts()) // 1
}
