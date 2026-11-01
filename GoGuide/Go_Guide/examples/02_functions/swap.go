// swap.go - Una funzione può restituire più risultati (slide 13-14).
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// Assegnazione multipla: entrambi i risultati vengono "spacchettati".
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
