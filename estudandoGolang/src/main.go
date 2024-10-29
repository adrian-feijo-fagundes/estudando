package main

import (
	"estudandoGolang/src/exercicios"
	"fmt"
)

// todo código fica aqui dentro
func main() {
	//aulas.Pacotes()
	//aulas.Tipos()
	//aulas.VarConst()
	//aulas.Funcao()
	//aulas.ZeroValues()
	// aulas.Listas()
	// aulas.Maps()
	//aulas.Structs()
	//sla := aulas.Data{Dia: 1, Mes: 1, Ano: 1}
	//fmt.Println(sla)
	// aulas.Condicionais()

	//aulas.ContagemRegressiva()

	// fmt.Println(exercicios.MakeNegative(0))
	// fmt.Println(exercicios.MakeNegative(-42))
	// fmt.Println(exercicios.MakeNegative(84))
	// fmt.Println(exercicios.Summation(8))
	// fmt.Println(exercicios.Summation(2))
	// fmt.Println(exercicios.Summation(8))
	// fmt.Println(exercicios.Summation(8))
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	var arr2 []int
	fmt.Println(exercicios.CountPositiveSumNegatives(arr))
	fmt.Println(exercicios.CountPositiveSumNegatives(arr2))
	fmt.Println(exercicios.StringToNumber("123"))

}
