package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"slices"
	"text/template"

	"github.com/goccy/go-yaml"
)

//go:embed manifest.yaml
var manifestData []byte

//go:embed workflows
var workflows embed.FS

type target struct {
	Bundles []string `yaml:"bundles"`
	Checks  []string `yaml:"checks"`
}

type manifest struct {
	Bundles  map[string][]string `yaml:"bundles"`
	Defaults struct {
		Checks []string `yaml:"checks"`
	} `yaml:"defaults"`
	Targets map[string]target `yaml:"targets"`
}

func main() {
	if len(os.Args) < 2 {
		panic("usage: workflow-sync <list|render> [args]")
	}

	m := load()

	switch os.Args[1] {
	case "list":
		list(m, slices.Contains(os.Args[2:], "--json"))
	case "render":
		if len(os.Args) < 4 {
			panic("usage: workflow-sync render <repo> <outdir>")
		}
		render(m, os.Args[2], os.Args[3])
	default:
		panic("unknown command: " + os.Args[1])
	}
}

func load() manifest {
	var m manifest
	if err := yaml.Unmarshal(manifestData, &m); err != nil {
		panic(err)
	}
	return m
}

func list(m manifest, asJSON bool) {
	repos := make([]string, 0, len(m.Targets))
	for repo := range m.Targets {
		repos = append(repos, repo)
	}
	slices.Sort(repos)

	if asJSON {
		b, err := json.Marshal(map[string][]string{"repo": repos})
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
		return
	}

	for _, repo := range repos {
		fmt.Println(repo)
	}
}

func render(m manifest, repo, out string) {
	t, ok := m.Targets[repo]
	if !ok {
		panic("unknown target: " + repo)
	}

	checks := t.Checks
	if len(checks) == 0 {
		checks = m.Defaults.Checks
	}
	data := map[string]any{"Checks": checks}

	dst := filepath.Join(out, ".github", "workflows")
	if err := os.MkdirAll(dst, 0o755); err != nil {
		panic(err)
	}

	for _, bundle := range t.Bundles {
		files, ok := m.Bundles[bundle]
		if !ok {
			panic("unknown bundle: " + bundle)
		}
		for _, file := range files {
			b, err := workflows.ReadFile(path.Join("workflows", file))
			if err != nil {
				panic(err)
			}
			tmpl, err := template.New(file).Delims("[[", "]]").Parse(string(b))
			if err != nil {
				panic(err)
			}
			f, err := os.Create(filepath.Join(dst, file))
			if err != nil {
				panic(err)
			}
			if err := tmpl.Execute(f, data); err != nil {
				panic(err)
			}
			f.Close()
			fmt.Fprintln(os.Stderr, "rendered", repo, "→", file)
		}
	}
}
