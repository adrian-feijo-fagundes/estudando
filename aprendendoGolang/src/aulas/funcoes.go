package aulas

import "fmt"

// Função começando com letra maiuscula:
// Função é PÚBLICA
// Ela pode ser utilizada fora do próprio pacote
// Como utilizaria ela fora do pacote: main.PrintaNome(nome)
func Funcao() {
	soma := soma(10, 12)
	sub := subtracao(soma, 1)
	soma2, sub2 := sumAndsub(10, 1)
	soma3, _ := sumAndsub(2, 1)

	fmt.Println("1ª SOMA: ", soma)
	fmt.Println("1ª SUBTRAÇÃO:  ", sub)
	fmt.Println("2ª SOMA: ", soma2)
	fmt.Println("2ª SUBTRAÇÃO: ", sub2)
	fmt.Println("3ª SOMA: ", soma3)
	fmt.Println(hello("Adrian"))

}

// Função começando com letra minúscula:
// Função é PRIVADA
// Ela só pode ser utilizada no próprio pacote
func soma(x int, y int) int {
	return x + y
}

func subtracao(x int, y int) int {
	return x - y
}

func sumAndsub(x, y int) (int, int) { // É possivel retornar mais de um valor
	return x + y, x - y
}
func hello(name string) string {
	return "Olá " + name + "!"
}
