package aulas

import "fmt"

func Tipos() {
	//bool (true/false)
	fmt.Printf("Type: %T - Value %v\n", false, false)

	// string - sequência de bytes
	fmt.Printf("Type: %T - Value %v\n", "Adrian", "Adrian")
	fmt.Printf("Type: %T - Value %v\n", "1", "1")

	// int
	fmt.Printf("Type: %T - Value %v\n", 1, 1)

	// float (float64/float32) - decimal
	fmt.Printf("Type: %T - Value %v\n", 1.233, 1.233)

}
