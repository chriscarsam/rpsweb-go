package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error when analyzing templates", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Message string
	}{
		Title:   "Home page",
		Message: "¡Welcome to rock, paper, scissors!",
	}

	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering templates", http.StatusInternalServerError)
		return
	}
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
