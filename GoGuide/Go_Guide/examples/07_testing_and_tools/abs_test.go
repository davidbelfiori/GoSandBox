package tools

import "testing"

// I test stanno in file *_test.go; le funzioni si chiamano TestXxx e
// ricevono un *testing.T.
func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		// Errorf segna il test come fallito ma lo lascia proseguire.
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

func TestAbsPositive(t *testing.T) {
	if got := Abs(5); got != 5 {
		t.Errorf("Abs(5) = %d; want 5", got)
	}
}

// Verifica che le due costruzioni producano lo stesso risultato.
func TestFillEquivalent(t *testing.T) {
	a, b := FillAppend(100), FillMake(100)
	if len(a) != len(b) {
		t.Fatalf("lunghezze diverse: %d vs %d", len(a), len(b))
	}
	for i := range a {
		if a[i] != b[i] {
			t.Fatalf("diverse in posizione %d", i)
		}
	}
}
