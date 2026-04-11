package brand

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:generate npx @tailwindcss/cli -i src/app.css -o dist/dbx.css --minify

//go:embed dist/*
var dist embed.FS

var Dist http.FileSystem

func init() {
	f, err := fs.Sub(dist, "dist")
	if err != nil {
		panic(err.Error())
	}
	Dist = http.FS(f)
}
