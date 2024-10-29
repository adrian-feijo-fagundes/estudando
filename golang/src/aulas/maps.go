package aulas

import "fmt"

func Maps() {
	fmt.Println("Maps")

	// 2 - Maps: Heterogêneos
	// pode misturar tipos
	// estrutura chave - valor
	// [key] = value
	// chave tem um tipo, e o valor pode ter outro
	// map[string]int
	// { "steph": 28, "bento": 4}
	// map[string]string
	// { "steph": "cardoso", "bento": "cardoso" }

	idade := map[string]int{}

	idade["adrian"] = 21
	idade["miqueias"] = 16

	fmt.Println(idade)
	fmt.Println(idade["adrian"])
	fmt.Println(idade["miqueias"])

	anoNasc := map[string]int{
		"adrian":   2002,
		"miqueias": 2008,
	}

	anoNasc["ester"] = 2007
	fmt.Println(anoNasc)
}
