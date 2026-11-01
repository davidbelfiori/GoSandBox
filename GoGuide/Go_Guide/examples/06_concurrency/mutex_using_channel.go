// mutex_using_channel.go - Esercizio: un mutex realizzato con un channel
// (slide 61-62).
//
// Idea: un channel con buffer di capacità 1 contiene al più un "gettone".
// Chi possiede il gettone detiene il lock; chi lo trova assente si blocca
// sulla ricezione.
package main

import (
	"fmt"
	"sync"
)

// Nota: "go vet" segnala "return copies lock value" su NewLock perché il
// tipo ha metodi Lock/Unlock e vet applica l'euristica dei lock non
// copiabili (come sync.Mutex). Qui la copia è innocua: la struct contiene
// solo un riferimento al channel, quindi tutte le copie condividono lo
// stesso "gettone". Il codice è quello della slide.
type Lock struct {
	ch chan bool
}

func NewLock() Lock {
	lock := Lock{make(chan bool, 1)} // channel con buffer
	lock.ch <- true                  // gettone iniziale: il lock parte libero
	return lock
}

// Lock prende il gettone; se non c'è, blocca finché qualcuno lo restituisce.
func (l *Lock) Lock() {
	<-l.ch // receive
}

// Unlock restituisce il gettone.
func (l *Lock) Unlock() {
	l.ch <- true // send
}

func main() {
	lock := NewLock()
	counter := 0

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				lock.Lock()
				counter++ // sezione critica
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("counter =", counter) // sempre 100000
}
