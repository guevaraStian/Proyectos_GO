package main

import (
	"fmt"
)

func main() {
	mimapa := map[string]string{
		"colombia":  "bogota",
		"venezuela": "caracas",
		"brasil":    "brasilia",
		"chile":     "santiago de chile",
	}

	fmt.Println("El mapeo de los paises ", mimapa)

	fmt.Println("la capital de venezuela es ", mimapa["venezuela"])

	mimapa["argentina"] = "buenos aires"

	fmt.Println("El mapa de paises es: ", mimapa["argentina"])

	delete(mimapa, "venezuela")
	fmt.Println("la capital de venezuela es ", mimapa["brasil"])

}
