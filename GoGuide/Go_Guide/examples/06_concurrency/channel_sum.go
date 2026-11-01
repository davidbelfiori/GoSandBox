// channel_sum.go - Somma distribuita tra due goroutine (slide 51).
//
// Pattern SPMD (Single Program Multiple Data): ogni goroutine esegue lo
// stesso codice su una porzione diversa dei dati.
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c) // prima metà: 7+2+8 = 17
	go sum(s[len(s)/2:], c) // seconda metà: -9+4+0 = -5

	// Le due ricezioni bloccano finché le goroutine non inviano.
	// Non si sa quale somma arrivi per prima, ma il totale è lo stesso
	// (l'addizione è commutativa).
	sum1, sum2 := <-c, <-c
	fmt.Println("Total sum:", sum1+sum2)
}
