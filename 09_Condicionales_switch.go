package main

import (
	"fmt"
)

func main() {
	var fruta string
	fmt.Print("Ingresar nombre de la fruta: ")
	fmt.Scan(&fruta)

	switch fruta {
	case "manzana":
		fmt.Println("Has ingresado manzana y cuesta 1000")
	case "pera":
		fmt.Println("Has ingresado pera y cuesta 2000")
	case "limon":
		fmt.Println("Has ingresado limon y cuesta 3000")
	case "naranja":
		fmt.Println("Has ingresado naranja y cuesta 4000")
	default:
		fmt.Println("No esta registrada esa fruta")

	}

}
