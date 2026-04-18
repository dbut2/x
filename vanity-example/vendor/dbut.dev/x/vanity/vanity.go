package vanity

import (
	"net/http"
	"runtime/debug"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

// Middleware returns a gin middleware that responds to ?go-get=1 probes with
// vanity import metadata pointing at repo, and otherwise lets the request pass
// through.
//
// The advertised import path comes from the running binary's main module
// (debug.ReadBuildInfo). Pass repo as the bare git URL ("github.com/owner/r")
// for modules at the repo root, or include the in-repo subdirectory
// ("github.com/owner/r/sub") for modules nested in the repo. With a subdir,
// the meta tag emits the 4-field go-import form so `go get` reads go.mod from
// the subdirectory of the repo.
func Middleware(repo string) gin.HandlerFunc {
	if repo == "" {
		panic("vanity: repo is empty")
	}
	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Path == "" {
		panic("vanity: cannot read main module from build info")
	}
	module := info.Main.Path

	bare, subdir := splitRepo(repo)
	bare = normalizeRepo(bare)
	browsePath := ""
	if subdir != "" {
		browsePath = "/" + subdir
	}

	return func(c *gin.Context) {
		if c.Query("go-get") != "1" {
			c.Next()
			return
		}
		v := &vanity{
			Module:     module,
			Vcs:        "git",
			ImportURL:  bare,
			Subdir:     subdir,
			BrowseRepo: bare,
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

// splitRepo separates a host/owner/name URL from any trailing in-repo subdir.
func splitRepo(s string) (repo, subdir string) {
	s = strings.TrimSuffix(s, "/")
	s = strings.TrimPrefix(s, "https://")
	s = strings.TrimPrefix(s, "http://")
	parts := strings.SplitN(s, "/", 4)
	if len(parts) < 4 {
		return s, ""
	}
	return strings.Join(parts[:3], "/"), parts[3]
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
<meta name="go-import" content="{{.Module}} {{.Vcs}} {{.ImportURL}}{{if .Subdir}} {{.Subdir}}{{end}}">
<meta name="go-source" content="{{.Module}} {{.BrowseRepo}}{{.BrowsePath}} {{.BrowseRepo}}/tree/HEAD{{.BrowsePath}}{/dir} {{.BrowseRepo}}/blob/HEAD{{.BrowsePath}}{/dir}/{file}#L{line}">
</head>
</html>
`))
