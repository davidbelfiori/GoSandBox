// variadic.go - Funzioni variadiche (slide 68).
package main

import "fmt"

// Dentro la funzione "nums" è una slice []int.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// Per passare una slice esistente si usa il suffisso "...".
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// Zero argomenti: nums è una slice nil e il totale è 0.
	sum()
}
