package exercicios

import "testing"

// exemplo de Teste
func TestMakeNegative(t *testing.T) {
	resultado := MakeNegative(3)
	esperado := -3

	if resultado != esperado {
		t.Errorf("Make Negative: esperado %d, obtido %d", esperado, resultado)
	}
}
