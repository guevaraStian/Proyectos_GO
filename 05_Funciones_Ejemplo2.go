package main

import (
	"fmt"
)

var funciones = map[string]func(int, int) int{
	"suma":           func(a int, b int) int { return a + b },
	"resta":          func(a int, b int) int { return a - b },
	"multiplicacion": func(a int, b int) int { return a * b },
	"division":       func(a int, b int) int { return a / b },
}

func presentarresultados(operacion string, a int, b int) {

	f, exists := funciones[operacion]

	if !exists {
		fmt.Println("Operacion no valida")
		return
	}

	fmt.Println("Para a= ", a, "y b = ", b, " la ", operacion, " resultantes es ", f(a, b))
}

func main() {
	presentarresultados("suma", 5, 8)
	presentarresultados("resta", 5, 8)
	presentarresultados("multiplicacion", 5, 8)
	presentarresultados("division", 5, 8)
	presentarresultados("novalido", 5, 8)
}
