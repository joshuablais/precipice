package main

import (
	"github.com/joshuablais/precipice/pages"
	"net/http"
)

func Routes(mux *http.ServeMux) {
	page := NewPageHandler()

	// Public routes
	mux.HandleFunc("GET /", page(pages.Index()))
	// mux.HandleFunc("GET /contact", page(public.ContactPage))
	// mux.HandleFunc("GET /marketing", page(public.MarketingPage))

	// Authentication routes
	// mux.HandleFunc("GET /register", page(authentication.SignupPage))
	mux.HandleFunc("GET /login", page(pages.LoginPage()))

	// // Blog routes
	// mux.HandleFunc("GET /blog", BlogIndexHandler(cfg))
	// mux.HandleFunc("GET /blog/{slug}", BlogPostHandler(cfg))

	// RSS feed
	// mux.HandleFunc("GET /index.xml", RSSHandler(cfg))
}
