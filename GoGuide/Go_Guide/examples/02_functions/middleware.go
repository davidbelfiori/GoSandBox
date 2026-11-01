// middleware.go - Closure come middleware (slide 40).
//
// La slide mostra timed(), che avvolge un handler HTTP per misurarne il tempo
// di elaborazione. Per rendere l'esempio eseguibile senza aprire porte di
// rete, la richiesta viene simulata con il package httptest.
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

// timed riceve un handler e ne restituisce uno nuovo che esegue codice
// "prima" e "dopo" la chiamata all'handler originale. f è catturato dalla
// closure.
func timed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		// Il tempo esatto cambia a ogni esecuzione: lo si arrotonda per
		// rendere l'output (quasi sempre) stabile.
		fmt.Println("The request took", time.Since(start).Round(50*time.Millisecond))
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond) // lavoro simulato
	fmt.Fprintln(w, "hello")
}

func main() {
	// Con un server vero si scriverebbe: http.HandleFunc("/hello", timed(hello))
	handler := timed(hello)

	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()
	handler(rec, req)

	fmt.Print("risposta: ", rec.Body.String())
}
