package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(10, 20)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido %d. Esperado: %d", total, 30)
	}
}
