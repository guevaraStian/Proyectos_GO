// Captura de pantalla del navegador hecho con el lenguaje de programacion GO
// go get -u github.com/chromedp/chromedp
// Se crea la imagen de la pantalla solicitada screenshot

package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://example.com"),
		chromedp.CaptureScreenshot(&buf),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("screenshot.png", buf, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Captura guardada como screenshot.png")
}
