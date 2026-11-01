// select_default.go - select non bloccante con default (slide 56, 59).
package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch3 := make(chan string, 1)

	// Ricezione non bloccante: nessun valore pronto -> scatta default.
	select {
	case msg := <-ch1:
		fmt.Println("ricevuto", msg)
	default:
		fmt.Println("nessun messaggio pronto")
	}

	// Invio non bloccante: ch3 ha un posto libero nel buffer.
	select {
	case ch3 <- "ciao":
		fmt.Println("inviato su ch3")
	default:
		fmt.Println("ch3 pieno")
	}

	// Secondo invio: il buffer è ora pieno, quindi si va in default.
	select {
	case ch3 <- "altro":
		fmt.Println("inviato su ch3")
	default:
		fmt.Println("ch3 pieno")
	}

	// Ora c'è un valore pronto su ch3: viene scelto il case di ricezione.
	select {
	case msg := <-ch3:
		fmt.Println("ricevuto", msg)
	default:
		fmt.Println("nessun messaggio pronto")
	}
}
