package aulas

import "fmt"

// Structs
// Forma de criar sua própria estrutura de dados
// Personalizar de acordo com a sua necessidade
// Podemos usar vários tipos diferentes

// type <nome da nossa estrutura> struct { <campos> }
type Data struct {
	Dia int
	Mes int
	Ano int
}
type Pessoa struct {
	nome    string
	anoNasc Data
	idade   int
}

type Profissao struct {
	Pessoa
	Tipo string
}

func Structs() {
	p1 := Pessoa{nome: "Adrian", idade: 21, anoNasc: Data{Dia: 1, Mes: 12, Ano: 2002}}
	p2 := Pessoa{nome: "Miqueias", idade: 16, anoNasc: Data{Dia: 7, Mes: 10, Ano: 2008}}

	fmt.Println(p1.nome)
	fmt.Println(p1.idade)

	//p1.Idade = 3
	// fmt.Println(p1.Idade)

	pessoas := []Pessoa{} // LISTA DE PESSOASS
	pessoas = append(pessoas, p1, p2)

	fmt.Println(pessoas)

	// structs + map
	alunos := map[string][]Pessoa{}
	alunos["programação"] = pessoas
	fmt.Println(alunos)

	// var alunos = map[string][]Pessoa{
	// 	"Programação": {{Nome: "Steph", Idade: 28}, {Nome: "Bento", Idade: 4}},
	// 	"Engenharia":  {{Nome: "Steph2", Idade: 28}, {Nome: "Bento2", Idade: 4}},
	// }
	// fmt.Println(alunos)

	// struct herdando campos de outra struct
	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.nome)
	fmt.Println(prof.Pessoa.idade)
	fmt.Println(prof.Tipo)
}
