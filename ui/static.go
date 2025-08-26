package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:generate templ generate
//go:generate npx @tailwindcss/cli -i app.css -o static/style.css

//go:embed static/*
var static embed.FS

var StaticContent http.FileSystem

func init() {
	f, err := fs.Sub(static, "static")
	if err != nil {
		panic(err.Error())
	}
	StaticContent = http.FS(f)
}
