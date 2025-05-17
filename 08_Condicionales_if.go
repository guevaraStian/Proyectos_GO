package main

import (
	"fmt"
)

func main() {
	var edad int = 19

	if edad >= 65 {
		fmt.Println("La persona esta jubilada")

	} else if edad >= 18 {
		fmt.Println("La persona es mayor de edad")

	} else {
		fmt.Println("La persona es menor de edad")
	}

	fmt.Println("Fin de programa")
}
