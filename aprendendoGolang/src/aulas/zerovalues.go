package aulas

import "fmt"

func ZeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("Inteiro %v\n", i) // 0
	fmt.Printf("Float %v\n", f)   // 0
	fmt.Printf("Bool %v\n", b)    // false
	fmt.Printf("String  %q\n", s) // ""
}
