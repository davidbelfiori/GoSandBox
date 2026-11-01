package tools

import "testing"

// sink è una variabile a livello di package: assegnarvi il risultato
// impedisce al compilatore di scartare o allocare sullo stack la slice
// (altrimenti, con inlining e dimensione costante, il benchmark di make
// risulterebbe senza allocazioni e il confronto sarebbe falsato).
var sink []int

// Benchmark: il runner sceglie b.N in modo adattivo, ripetendo il corpo
// finché la misura è stabile.
//
//	go test -bench=. ./07_testing_and_tools/

func BenchmarkFillAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = FillAppend(1000)
	}
}

func BenchmarkFillMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = FillMake(1000)
	}
}
