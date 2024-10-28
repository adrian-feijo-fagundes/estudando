package aulas

import "fmt"

func Condicionais() {

	idade := 16
	maioridade := idade > 17
	carteira := false
	podeDirigir := maioridade && carteira
	altura := 1.80

	soma := float64(8) + (altura)
	fmt.Println(soma)
	/* OPERADORES LÓGICOS
	&&	E
	||  OU
	!	NÃO
	*/
	/* COMPARAÇÃO */
	a := 10
	b := 20

	fmt.Println("Igualdade:", a == b)        // false
	fmt.Println("Desigualdade:", a != b)     // true
	fmt.Println("Maior que:", a > b)         // false
	fmt.Println("Menor que:", a < b)         // true
	fmt.Println("Maior ou igual a:", a >= b) // false
	fmt.Println("Menor ou igual a:", a <= b) // true

	/* MATEMATICOS
	Soma: 1 + 1
	Subtração 2 - 1
	Divisão: 10 / 2
	Multiplicação: 2 * 2
	Resto da divisão por x: 7%2 (resto da divisao por 2)
	*/

	if podeDirigir {
		fmt.Println("Voce pode dirigir")
	} else if maioridade {
		fmt.Println("Voce é de maior, mas não possui carteira")
	} else {
		fmt.Println("Voce ainda não é de maior")
	}

	// SWITCH CASE
	time := "gremio"
	switch time {
	case "inter":
		fmt.Println("INTER")
	case "gremio":
		fmt.Println("GREMIO")
	default:
		fmt.Println("Não conheço o time")

	}

}
