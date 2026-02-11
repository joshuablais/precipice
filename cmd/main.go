// Precipice is a UI Framework to end all frameworks
package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	// Static files
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	Routes(mux) // This goes LAST

	// Routes
	slog.Info("starting server", "addr", ":3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		slog.Error("server error", "err", err)
		os.Exit(1)
	}
}
