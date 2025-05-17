package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Contraseña puede ser en texo o en hash
	objetivo := "contra123"

	// Abrir archivo del diccionario con la librerias os
	archivo, err := os.Open("diccionario.txt")
	// Si no se pudo abrir el diccionario se muestra un error
	if err != nil {
		fmt.Println("Error al abrir el diccionario:", err)
		return
	}
	defer archivo.Close()

	// Con estos codigos se leer cada linea del diccionario y se compara
	scanner := bufio.NewScanner(archivo)
	intento := 0
	for scanner.Scan() {
		intento++
		palabra := strings.TrimSpace(scanner.Text())
		if palabra == objetivo {
			fmt.Printf("Se encontro la contraseña: %s (En el intento: %d)\n", palabra, intento)
			return
		}
	}

	// Si no se encontró la contraseña se muestra en siguiente mensaje
	fmt.Println("En el diccionario no se encontro similitudes.")
}
