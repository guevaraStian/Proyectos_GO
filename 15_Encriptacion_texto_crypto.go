// En este codigo se muestra como encriptar un texto en md5 usando
// la libreria crypto y su seccion md5
package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// Se crea la variable del texto y se le ingresa la informacion
	texto := ""
	fmt.Println("Por favor Ingrese texto a encriptar : ")
	fmt.Scan(&texto)
	// Se procede a usar la funcion md5 de la libreria
	hashmd5 := md5.Sum([]byte(texto))
	// la encriptacion se vuelve en digitos hexadecimales para mostrar en pantalla
	hashHex := hex.EncodeToString(hashmd5[:])

	hashsha256 := sha256.Sum256([]byte(texto))
	// la encriptacion se vuelve en digitos hexadecimales para mostrar en pantalla
	hashHex2 := hex.EncodeToString(hashsha256[:])

	fmt.Println("El texto original:", texto)
	fmt.Println("El MD5 hash:", hashHex)
	fmt.Println("El SHA256 hash:", hashHex2)
}
