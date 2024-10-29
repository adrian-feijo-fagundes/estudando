package exercicios

import (
	"fmt"
	"strconv"
	"strings"
)

// Exercicio 1
// transforma números positivos em negativos
func MakeNegative(x int) int {
	if x > 0 {
		return -x
	}
	return x
}

// Exercicio 2
func Summation(n int) int {
	var sum int

	for i := n; i > 0; i-- {
		sum += i
	}
	return sum
}

// Exercicio 3
// recebe uma lista e soma apenas os numeros positivos
func PositiveSum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		if num > 0 {
			sum += num
		}
	}

	return sum
}

// Exercicio 4
// conta os positivos e soma os negativos
func CountPositiveSumNegatives(numbers []int) []int {
	var positiveCount, negativeSum int
	if len(numbers) == 0 {
		return []int{}
	}
	for _, num := range numbers {
		if num > 0 {
			positiveCount++
		}
		if num < 0 {
			negativeSum += num
		}
	}

	return []int{positiveCount, negativeSum}
}

// Exercicio 5
func StringToNumber(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

// Exercicio 6
func NumberToString(n int) string {
	//return strconv.Itoa(n)
	return fmt.Sprintf("%d", n)
}

// Exercicio 7
func RepeatStr(repetitions int, value string) (result string) {
	for i := 0; i < repetitions; i++ {
		result += value
	}
	// strings.Repeat(value, repititions)
	return result
}

// Exercicio 8
// inverte positivo para negativo e vice-versa
func Invert(arr []int) []int {
	res := []int{}
	if len(arr) == 0 {
		return res
	}
	for _, num := range arr {
		res = append(res, (num * -1))
	}
	return res
}

// Exercicio 9
// par ou impar
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

// Exercicio 10
// anos dos animais Humano, Gato, Cachorro
func CalculateYears(years int) (result [3]int) {
	switch years {
	case 1:
		result = [3]int{1, 15, 15}
	case 2:
		result = [3]int{2, 24, 24}
	default:
		result = [3]int{years, 24 + 4*(years-2), 24 + 5*(years-2)}
	}
	return
}

// Exercicio 11
// Pedra Papel Tesoura
func Rps(p1, p2 string) string {
	var m = map[string]string{
		"rock":     "paper",
		"paper":    "scissors",
		"scissors": "rock",
	}

	if p1 == p2 {
		return "Draw!"
	}
	if m[p1] == p2 {
		return "Player 2 won!"
	}
	return "Player 1 won!"
}

// Exercicio 12
func Well(x []string) string {
	goodIdeas := 0

	for _, idea := range x {
		if idea == "good" {
			goodIdeas++
		}
	}
	switch goodIdeas {
	case 0:
		return "Fail!"
	case 1, 2:
		return "Publish!"
	}
	return "I smell a series!"
}

// Exercicio 13
func Points(games []string) int {
	// points := 0
	const WIN_POINTS = 3
	const DRAW_POINTS = 1
	// for _, game := range games {
	// 	if game[0] > game[2] {
	// 		points += 3
	// 	}
	// 	if game[0] == game[2] {
	// 		points++
	// 	}
	// }

	// return points
	result := 0
	for _, game := range games {
		str := strings.Split(game, ":")
		x, y := str[0], str[1]
		switch {
		case x > y:
			result += WIN_POINTS
		case x == y:
			result += DRAW_POINTS
		}
	}
	return result
}

// exercicio 14
func PartList(words []string) string {

	var combined string

	for i := 1; i < len(words); i++ {
		word := fmt.Sprintf("(%s, %s)", strings.Join(words[:i], " "), strings.Join(words[i:], " "))
		combined += word
	}

	return combined
}

// Exercicio 15
func SortMyString(word string) string {
	var a, b string
	for i, letter := range word {
		if i%2 == 0 {
			a += string(letter)
		} else {
			b += string(letter)
		}
	}
	return a + " " + b
}
