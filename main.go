package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {
	// Create router
	router := http.NewServeMux()

	// Configure routes
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	log.Printf("Sevidor listening in http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))

}
