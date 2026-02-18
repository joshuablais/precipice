// Precipice is a UI Framework to end all frameworks
package main

import (
	"github.com/joshuablais/precipice/internal/config"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.LoadConfig()
	mux := http.NewServeMux()

	// Static file serving
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	Routes(mux, cfg)

	// Routes
	slog.Info("starting server", "addr", ":3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		slog.Error("server error", "err", err)
		os.Exit(1)
	}
}
