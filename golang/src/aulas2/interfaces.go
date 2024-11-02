package aulas2

import "fmt"

type Animal interface {
	Barulho() string
}

type Cachorro struct {
	Raca  string
	Cor   string
	Patas int
}

type Gato struct {
	Cor   string
	Patas int
}

func (c Cachorro) Barulho() string {
	return "au au"
}

func (g Gato) Barulho() string {
	return "miau"
}

func fazBarulho(animal Animal) string {
	return animal.Barulho()
}

func Interfaces() {
	cinzinha := Gato{Cor: "cinza", Patas: 4}
	bolinha := Cachorro{Raca: "vira-lata", Cor: "preto", Patas: 4}

	fmt.Println(fazBarulho(cinzinha))
	fmt.Println(fazBarulho(bolinha))
}
