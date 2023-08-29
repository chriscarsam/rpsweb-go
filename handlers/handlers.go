package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home page")
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create new game")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Game")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Play")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About")
}
