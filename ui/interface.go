package ui

import (
	"html/template"
	"net/http"
)

type Handler struct{}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("ui/html/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
