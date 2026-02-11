package main

import (
	"github.com/a-h/templ"
	"net/http"
)

func NewPageHandler() func(templ.Component) http.HandlerFunc {
	return func(component templ.Component) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			component.Render(r.Context(), w)
		}
	}
}
