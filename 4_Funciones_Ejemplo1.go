package main

import "fmt"

func suma(a int, b int) int {
	return a + b
}
func resta(a int, b int) int {
	return a - b
}
func multiplicacion(a int, b int) int {
	return a * b
}
func division(a int, b int) int {
	return a / b
}

func presentarresultadossuma(a int, b int) {
	fmt.Println("Para a= ", a, "y b = ", b, " la suma resultantes es ", suma(a, b))
}
func presentarresultadosresta(a int, b int) {
	fmt.Println("Para a= ", a, "y b = ", b, " la resta resultantes es ", resta(a, b))
}
func presentarresultadosmultiplicacion(a int, b int) {
	fmt.Println("Para a= ", a, "y b = ", b, " la multiplicacion resultantes es ", multiplicacion(a, b))
}
func presentarresultadosdivision(a int, b int) {
	fmt.Println("Para a= ", a, "y b = ", b, " la division resultantes es ", division(a, b))
}

func main() {
	presentarresultadossuma(5, 8)
	presentarresultadossuma(10, 12)

	presentarresultadosresta(78, 23)
	presentarresultadosresta(7, 2)

	presentarresultadosmultiplicacion(7, 2)
	presentarresultadosmultiplicacion(7, 2)

	presentarresultadosdivision(7, 2)
	presentarresultadosdivision(7, 2)

}
