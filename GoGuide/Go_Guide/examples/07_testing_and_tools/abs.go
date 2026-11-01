// Package tools contiene la funzione sotto test per gli esempi di
// testing e benchmarking (slide 71-72).
//
// Non è un package main: non si esegue con "go run", ma con
//
//	go test ./07_testing_and_tools/
//	go test -bench=. ./07_testing_and_tools/
package tools

// Abs restituisce il valore assoluto di n.
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// FillAppend costruisce una slice di n elementi partendo da una slice
// vuota e usando append (può dover riallocare più volte).
func FillAppend(n int) []int {
	var s []int
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	return s
}

// FillMake costruisce la stessa slice allocando in anticipo la capacità
// con make: una sola allocazione.
func FillMake(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	return s
}
