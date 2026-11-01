// range2.go - range: ignorare indice o valore, iterare su map (slide 36).
package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	// "_" (blank identifier) scarta l'indice: Go vieta variabili inutilizzate.
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Su una map range produce chiave e valore. L'ordine di iterazione NON
	// è specificato (in Go è volutamente casuale): un ciclo diretto
	// "for k, v := range kvs" stampa le due coppie in ordine variabile.
	// Per avere un output stabile si itera su un elenco di chiavi
	// ordinato a mano e si consulta la map.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for _, k := range []string{"a", "b"} {
		fmt.Printf("%s -> %s\n", k, kvs[k])
	}

	// Con la sola chiave si omette il secondo valore.
	n := 0
	for k := range kvs {
		_ = k
		n++
	}
	fmt.Println("chiavi iterate:", n)

	// Anche le stringhe: range restituisce indice di byte e rune (Unicode).
	for i, r := range "gò" {
		fmt.Printf("%d:%c ", i, r)
	}
	fmt.Println()
}
