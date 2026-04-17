package vanity

import (
	"net/http"
	"runtime/debug"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

// Middleware returns a gin middleware that responds to ?go-get=1 probes with
// the vanity import metadata pointing at repo, and otherwise lets the request
// pass through to the rest of the chain. The advertised import path is the
// main module path read from the running binary's build info — so define it
// however you like in go.mod.
func Middleware(repo string) gin.HandlerFunc {
	if repo == "" {
		panic("vanity: repo is empty")
	}
	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Path == "" {
		panic("vanity: cannot read main module from build info")
	}
	v := &vanity{
		Module: info.Main.Path,
		Repo:   normalizeRepo(repo),
	}
	return func(c *gin.Context) {
		if c.Query("go-get") == "1" {
			v.serve(c.Writer)
			c.Abort()
			return
		}
		c.Next()
	}
}

type vanity struct {
	Module string
	Repo   string
}

func normalizeRepo(repo string) string {
	repo = strings.TrimRight(repo, "/")
	if !strings.Contains(repo, "://") {
		repo = "https://" + repo
	}
	return repo
}

func (v *vanity) serve(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = tmpl.Execute(w, v)
}

var tmpl = template.Must(template.New("vanity").Parse(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="go-import" content="{{.Module}} git {{.Repo}}">
<meta name="go-source" content="{{.Module}} {{.Repo}} {{.Repo}}/tree/HEAD{/dir} {{.Repo}}/blob/HEAD{/dir}/{file}#L{line}">
</head>
</html>
`))
