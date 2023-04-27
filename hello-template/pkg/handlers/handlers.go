package handlers

import (
	"net/http"

	"github.com/huydoan/go-course/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html.tmpl")

}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html.tmpl")
}
