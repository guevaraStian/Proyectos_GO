package main

import (
	"fmt"
)

func MostrarTexto(name string) {
	fmt.Println("Cual es tu nombre desde la funcion: ", name)

}
func ObtenerAño() int {
	var edad int
	fmt.Println("Cual es tu edad: ", edad)
	fmt.Scan(&edad)
	return edad
}
func main() {
	var name string
	fmt.Println("Cual es tu nombre: ")
	fmt.Scan(&name)
	fmt.Println("Tu nombre es ", name)
	fmt.Println("Tu edad es ", ObtenerAño())
	MostrarTexto(name)

}
