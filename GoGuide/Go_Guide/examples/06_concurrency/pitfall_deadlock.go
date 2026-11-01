// pitfall_deadlock.go - ERRORE COMUNE: deadlock su channel senza buffer.
//
// ATTENZIONE: questo programma TERMINA VOLUTAMENTE con un errore fatale:
//
//	fatal error: all goroutines are asleep - deadlock!
//
// L'invio su un channel senza buffer blocca finché qualcuno riceve, ma
// l'unica goroutine esistente è quella bloccata: nessuno riceverà mai.
// Il runtime lo rileva e abortisce il programma (exit code 2).
package main

import "fmt"

func main() {
	ch := make(chan int)

	// ERRORE: invio bloccante nella stessa goroutine che dovrebbe ricevere.
	ch <- 1
	fmt.Println(<-ch) // mai raggiunto

	// Correzioni possibili: (1) channel con buffer make(chan int, 1);
	// (2) inviare da un'altra goroutine: go func() { ch <- 1 }().
}
