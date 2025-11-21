package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.Aboiut)

	port := ":8080"
	log.Printf("Servidor escuchado en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
