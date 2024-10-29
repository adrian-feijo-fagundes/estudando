package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// todo código fica aqui dentro
func main() {
	fmt.Println("Main")
	e := echo.New()

	fmt.Println(e)
}
