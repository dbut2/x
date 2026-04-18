package vanity

import (
	"net/http"
	"runtime/debug"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

func Middleware(repo string) gin.HandlerFunc {
	if repo == "" {
		panic("vanity: repo is empty")
	}
	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Path == "" {
		panic("vanity: cannot read main module from build info")
	}
	module := info.Main.Path

	repo, dir := splitRepo(repo)
	browsePath := ""
	if dir != "" {
		browsePath = "/" + dir
	}

	return func(c *gin.Context) {
		if c.Query("go-get") != "1" {
			c.Next()
			return
		}
		v := &vanity{
			Module:     module,
			Vcs:        "git",
			ImportURL:  repo,
			Subdir:     dir,
			BrowseRepo: repo,
			BrowsePath: browsePath,
		}
		v.serve(c.Writer)
		c.Abort()
	}
}

type vanity struct {
	Module     string
	Vcs        string
	ImportURL  string
	Subdir     string
	BrowseRepo string
	BrowsePath string
}

func splitRepo(s string) (repo, dir string) {
	parts := strings.SplitN(s, "/", 4)
	if len(parts) < 4 {
		return s, ""
	}
	return strings.Join(parts[:3], "/"), parts[3]
}

func (v *vanity) serve(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = tmpl.Execute(w, v)
}

var tmpl = template.Must(template.New("vanity").Parse(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="go-import" content="{{.Module}} {{.Vcs}} {{.ImportURL}}{{if .Subdir}} {{.Subdir}}{{end}}">
<meta name="go-source" content="{{.Module}} {{.BrowseRepo}}{{.BrowsePath}} {{.BrowseRepo}}/tree/HEAD{{.BrowsePath}}{/dir} {{.BrowseRepo}}/blob/HEAD{{.BrowsePath}}{/dir}/{file}#L{line}">
</head>
</html>
`))
