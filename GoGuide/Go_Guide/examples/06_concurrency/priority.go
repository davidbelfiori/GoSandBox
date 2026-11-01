// priority.go - Simulare priorità con più channel (slide 65).
//
// Lo scheduler di Go non ha priorità: la si implementa a livello di
// applicazione. Qui due channel (alta/bassa priorità); il consumatore
// controlla SEMPRE prima il channel ad alta priorità con un select non
// bloccante e passa a quello a bassa priorità solo se il primo è vuoto.
package main

import "fmt"

func main() {
	high := make(chan string, 3)
	low := make(chan string, 3)

	// I task a bassa priorità vengono inseriti per primi, di proposito.
	low <- "low-1"
	low <- "low-2"
	high <- "high-1"
	high <- "high-2"

	const total = 4
	for done := 0; done < total; done++ {
		select {
		case t := <-high: // controllato per primo
			fmt.Println("eseguo", t)
		default:
			// Nessun task ad alta priorità: si serve il basso.
			select {
			case t := <-low:
				fmt.Println("eseguo", t)
			default:
			}
		}
	}
}
