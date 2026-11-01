// pitfall_send_closed.go - ERRORE COMUNE: invio su un channel chiuso
// (slide 54).
//
// ATTENZIONE: questo programma TERMINA VOLUTAMENTE con un panic:
//
//	panic: send on closed channel
//
// Si noti che la RICEZIONE da un channel chiuso è invece lecita.
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	close(ch)

	v, ok := <-ch // ricezione dal buffer residuo: 1 true
	fmt.Println(v, ok)
	v, ok = <-ch // channel chiuso e vuoto: 0 false
	fmt.Println(v, ok)

	fmt.Println("invio su channel chiuso...")
	ch <- 2 // PANIC
}
