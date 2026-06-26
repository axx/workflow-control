package main

import (
	"log/slog"
	"net/http"
	"workflow-control/ui"
)

func main() {
	uiHandler := &ui.Handler{}

	http.HandleFunc("/", uiHandler.Home)

	slog.Info("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
