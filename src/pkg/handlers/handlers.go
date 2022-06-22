package main

import (
	"net/http"

	"github.com/zahidhasanpapon/youTube-playlist-len.com/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, "home.page.tmpl")
}
