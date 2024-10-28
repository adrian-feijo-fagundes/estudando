package aulas

import "fmt"

func VarConst() {
	// var nome string
	// nome = "Adrian"
	// nome = "Mateus"
	// fmt.Println("Váriavel do tipo string ", nome)

	// var idade int
	// idade = 21
	// idade = 14
	// fmt.Println("Váriavel do tipo int ", idade)

	// Inferencia de tipo
	// var nome2 = "Miqueias"
	// fmt.Println(nome2)

	// var idade2 = 10
	// fmt.Println(idade2)

	// declarando varias variaveis
	// var x, y int = 1, 2
	// fmt.Println("x:", x, "y:", y)

	// forma rápida de declarar quando tem inferencia de tipo
	z := 3
	fmt.Println("z", z)

	// CONSTANTES
	const meuAnoDeNacimento = 2002
	fmt.Println(meuAnoDeNacimento)
}
