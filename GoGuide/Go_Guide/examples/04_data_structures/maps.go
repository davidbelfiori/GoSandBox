// maps.go - Map: tabelle hash chiave-valore (slide 34).
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	// Una map va creata con make (o con un letterale): la map nil
	// permette letture ma non scritture.
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967} // inserimento
	fmt.Println(m["Bell Labs"])                  // lettura

	m["Bell Labs"] = Vertex{1, 2} // aggiornamento
	fmt.Println(m)

	// Chiave inesistente: si ottiene il valore zero, senza errore.
	fmt.Println("chiave assente:", m["Google"])

	// Idioma "comma ok": distingue "assente" da "presente con valore zero".
	elem, ok := m["Google"]
	fmt.Println(elem, ok) // {0 0} false
	elem, ok = m["Bell Labs"]
	fmt.Println(elem, ok) // {1 2} true

	delete(m, "Bell Labs") // cancellazione (nessun errore se assente)
	fmt.Println(len(m))

	// Errore comune: scrivere su una map nil provoca panic.
	var nilMap map[string]int
	fmt.Println(nilMap["x"]) // lettura ok: 0
	defer func() { fmt.Println("recuperato:", recover()) }()
	nilMap["x"] = 1 // panic: assignment to entry in nil map
}
