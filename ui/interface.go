package ui

import (
	"html/template"
	"log/slog"
	"net/http"
)

type Handler struct{}

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	slog.Info("Handling request for home page")

	tmpl, err := template.ParseFiles("ui/html/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func (h *Handler) RegisterRoutes() *http.ServeMux {
	slog.Info("Registering UI routes")
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.home)
	return mux
}

func (h *Handler) RedirectToUI(w http.ResponseWriter, r *http.Request) {
	slog.Info("Redirecting to /ui/")
	http.Redirect(w, r, "/ui/", http.StatusFound)
}
