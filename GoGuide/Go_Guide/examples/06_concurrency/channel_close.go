// channel_close.go - close e rilevamento della chiusura (slide 54).
package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Il consumatore legge finché il channel è chiuso E vuoto.
	go func() {
		for {
			j, more := <-jobs // more == false: chiuso e senza altri valori
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Il produttore (il sender) invia e poi chiude: solo il sender
	// dovrebbe chiudere un channel.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done // attende il consumatore

	// Ricevere da un channel chiuso e vuoto restituisce subito il valore
	// zero e ok == false.
	v, ok := <-jobs
	fmt.Println(v, ok) // 0 false
}
