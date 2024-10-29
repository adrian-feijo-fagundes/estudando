package aulas

import "fmt"

func Listas() {
	// LISTAS

	// 1 - Arrays e Slices: Homogêneos
	// todos os elementos tem o mesmo tipo
	// [1, 2, 3, 4, 5, 6] - []int
	// ["steph", "bento", "golang"] - []string

	// Array

	// Tamanho fixo, de zero ou mais elementos do mesmo tipo
	// Acessamos os valores com índice: a[0], a[1]...
	// Função embutida len() retorna o tamanho do array
	// Por conta do tamanho fixo, não é tão usado. Só em casos específicos
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"

	fmt.Println(array[0], array[1])
	fmt.Println("tamanho do array:", len(array))

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[0:4])
	fmt.Println(numPrimos[0:2]) // "printa" a partir da posicao 0...até a posicao anterior a 2

	// Slice

	// Tipo o array, mas sem tamanho fixo
	// Acessamos os valores com índice: a[0], a[1]...
	// Função embutida len() retorna o tamanho do slice
	// Função append() usada para adicionar valores nesse slice
	var slice []string

	slice = append(slice, "oi")
	slice = append(slice, "meu")
	slice = append(slice, "querido")
	slice = append(slice, "amigo")

	fmt.Println(slice)

}
