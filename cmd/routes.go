package main

import (
	"github.com/joshuablais/precipice/internal/config"
	"github.com/joshuablais/precipice/pages"
	"net/http"
)

func Routes(mux *http.ServeMux, cfg *config.Config) {
	page := NewPageHandler(cfg)

	// Public routes
	mux.HandleFunc("GET /", page(pages.IndexPage(cfg)))
	mux.HandleFunc("GET /components", page(pages.ComponentsPage(cfg)))

	// TODO
	mux.HandleFunc("GET /buttons", page(pages.ButtonsPage(cfg)))
	// mux.HandleFunc("GET /fields", page(pages.Fields()))

	// Authentication routes
	// mux.HandleFunc("GET /register", page(authentication.SignupPage))
	mux.HandleFunc("GET /login", page(pages.LoginPage(cfg)))

	// // Blog routes
	// mux.HandleFunc("GET /blog", BlogIndexHandler(cfg))
	// mux.HandleFunc("GET /blog/{slug}", BlogPostHandler(cfg))

	// RSS feed
	// mux.HandleFunc("GET /index.xml", RSSHandler(cfg))
}
