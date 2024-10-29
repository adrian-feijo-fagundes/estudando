package aulas

import (
	"fmt"
	"time"
)

func Loops() {
	for i := 0; i < 10; i++ {
		fmt.Println("valor do i: ", i)
	}
}
func LoopCondicional() {
	contador := 0
	for contador < 10 {
		fmt.Println("valor do contador: ", contador)
		contador++
	}
}
func LoopInfinito() {
	i := 0
	for {
		if i >= 5 {
			break // Sai do loop quando i é maior ou igual a 5
		}
		fmt.Println("valor do i do loop infinito:", i)
		i++
	}
}
func LoopSlice() {
	frutas := []string{"laranja", "maça", "banana", "uva", "kiwi"}
	for _, fruta := range frutas {
		fmt.Println("Fruta", fruta)
		//time.Sleep(1 * time.Second)
	}
	for index, fruta := range frutas {
		fmt.Println(index, "Fruta", fruta)
	}
}
func LoopMap() {
	idades := map[string]int{
		"Adrian":   21,
		"Miqueias": 16,
		"Celio":    43,
		"Joseane":  42,
	}

	for nome, idade := range idades {
		fmt.Println(nome, idade)
	}
}
func LoopContinue() {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue // Ignora números pares
		}
		fmt.Println(i) // Apenas números ímpares serão impressos
	}
}
func LoopString() {
	str := "golang"

	for index, char := range str {
		fmt.Printf("Índice %d: %c\n", index, char)
	}
}

func ContagemRegressiva() {
	i := 10
	for {
		if i == 0 {
			break // Sai do loop quando i é maior ou igual a 5
		}
		fmt.Println(i)
		time.Sleep(1 * time.Second)
		i--
	}
}
