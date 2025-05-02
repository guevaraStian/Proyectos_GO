// En este codigo se muestra como crear una contraseña robusta
// Usando la libreria crypto se crea la contraseña
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Se crea la funcion generar contraseña con la descripcion de los caracteres
func generarContrasena(longitud int) (string, error) {
	const caracteres = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	contrasena := make([]byte, longitud)

	// con este for escogemos aleatoriamente cada digito hasta que se cumpla la cantidad de caracteres
	for i := 0; i < longitud; i++ {
		numAleatorio, err := rand.Int(rand.Reader, big.NewInt(int64(len(caracteres))))
		if err != nil {
			return "", err
		}
		contrasena[i] = caracteres[numAleatorio.Int64()]
	}

	return string(contrasena), nil
}

func main() {
	var longitud int
	fmt.Print("Ingresa la longitud de la contraseña: ")
	_, err := fmt.Scan(&longitud)
	if err != nil || longitud <= 0 {
		fmt.Println("Longitud no válida.")
		return
	}

	// Si sale un error se muestra en pantalla el siguiente mensaje o de lo contrario
	// se muestra la contraseña encriptada
	contrasena, err := generarContrasena(longitud)
	if err != nil {
		fmt.Println("Error al generar la contraseña:", err)
		return
	}

	fmt.Println("Contraseña segura generada:", contrasena)
}
