package handlers

import (
	"net/http"

	"github.com/KiroLakestrike/bedAndBreakfast/pkg/render"
)

// Handler for our Home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// Handler for our About page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
