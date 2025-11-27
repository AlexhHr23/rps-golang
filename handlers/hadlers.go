package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, "Error al analizar plantillas", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Message string
	}{
		Title:   "Pagina de inicio",
		Message: "Bienvenid a Piedra, Papel o Tijera",
	}

	err = tpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Crear nuevo juevo")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Juego")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Jugar")
}

func Aboiut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Acerca de")
}
