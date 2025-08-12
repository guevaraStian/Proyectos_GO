// Dashborard donde le muestran datos en una grafica con GO
// Se importan las librerias
package main

import (
	"html/template"
	"log"
	"net/http"
)

// Se crea la funcion main con la direccion del archivo javascript para la visualizacion dinamica
func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", dashboardHandler)

	// Se habilita la pagina con la grafica
	log.Println("Servidor corriendo en http://localhost:1234")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Datos simulados (pueden venir de una base de datos)
	data := map[string]interface{}{
		"Title":    "Dashboard Simple",
		"Users":    111,
		"Visitors": 222,
	}

	tmpl.Execute(w, data)
}
