// Informacion del navegador hecho con el lenguaje de programacion GO
// go get -u github.com/chromedp/chromedp
// Se solicita informacion de los datos del navegador
// Se solicita el titulo de la pagina web y se imprime en pantalla

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// Crear contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Crear instancia de Chrome
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// Resultado del título
	var title string

	// Ejecutar tareas
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://golang.org"),
		chromedp.Title(&title),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("El título de la página web es el siguiente:", title)
}
