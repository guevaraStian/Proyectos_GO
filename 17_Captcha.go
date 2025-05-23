// Captcha hecho con el lenguaje de programacion GO
// go get github.com/dchest/captcha
// Visita http://localhost:8080/generar_codigo_captcha para obtener un CAPTCHA ID.
// Ve a http://localhost:8080/captcha/<ID>.png para ver la imagen del CAPTCHA.
// Ingresa el texto que ves en la imagen y verifícalo en http://localhost:8080/verificar?id=<ID>&value=<texto>.

package main

import (
	"net/http"

	"github.com/dchest/captcha"
)

func main() {
	// Con el siguiente codigo se inicializa el captcha
	http.Handle("/captcha/", captcha.Server(captcha.StdWidth, captcha.StdHeight))

	// Luego se crea el codigo ID que esta relacionado a la imagen que hay
	http.HandleFunc("/generar_codigo_captcha", func(w http.ResponseWriter, r *http.Request) {
		captchaId := captcha.New()
		w.Write([]byte(captchaId)) // Se muestra el ID del captcha
	})

	// En este codigo se muestra una validacion del codigo en la imagen
	http.HandleFunc("/verificar", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		value := r.URL.Query().Get("value")

		if captcha.VerifyString(id, value) {
			w.Write([]byte("✅ CAPTCHA correcto"))
		} else {
			w.Write([]byte("❌ CAPTCHA incorrecto"))
		}
	})

	// Se activa la ruta http por el puerto 8080
	http.ListenAndServe(":8080", nil)
}
