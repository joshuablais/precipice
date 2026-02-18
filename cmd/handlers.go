package main

import (
	"github.com/a-h/templ"
	"github.com/joshuablais/precipice/internal/config"
	"net/http"
)

func NewPageHandler(cfg *config.Config) func(templ.Component) http.HandlerFunc {
	return func(component templ.Component) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			component.Render(r.Context(), w)
		}
	}
}
