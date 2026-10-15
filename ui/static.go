package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:generate go tool templ generate

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
