package brand

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:generate cp src/plain.css ../hugo/static/plain.css

//go:embed src/*.css
var src embed.FS

var Dist http.FileSystem

func init() {
	f, err := fs.Sub(src, "src")
	if err != nil {
		panic(err.Error())
	}
	Dist = http.FS(f)
}
