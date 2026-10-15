package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/goccy/go-yaml"
)

//go:embed manifest.yaml
var manifestData []byte

//go:embed workflows
var workflows embed.FS

const marker = "# DO NOT EDIT. Synced from dbut2/x\n"

const header = marker + "# https://github.com/dbut2/x/blob/main/workflow-sync/workflows/%s\n\n"

type target struct {
	Bundles []string `yaml:"bundles"`
	Checks  []string `yaml:"checks"`
	Gomod   string   `yaml:"gomod"`
	RunsOn  string   `yaml:"runs-on"`
}

type bundle struct {
	Files  []string `yaml:"files"`
	Checks []string `yaml:"checks"`
}

type manifest struct {
	Bundles map[string]bundle `yaml:"bundles"`
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

	bundles := make([]bundle, 0, len(t.Bundles))
	for _, name := range t.Bundles {
		b, ok := m.Bundles[name]
		if !ok {
			panic("unknown bundle: " + name)
		}
		bundles = append(bundles, b)
	}

	var checks []string
	for _, b := range bundles {
		checks = append(checks, b.Checks...)
	}
	for _, c := range t.Checks {
		if !slices.Contains(checks, c) {
			checks = append(checks, c)
		}
	}
	runsOn := t.RunsOn
	if runsOn == "" {
		runsOn = "ubuntu-latest"
	}
	gomod := t.Gomod
	if gomod == "" {
		gomod = "go.mod"
	}
	data := map[string]any{"Checks": checks, "Gomod": gomod, "GomodDir": path.Dir(gomod), "RunsOn": runsOn}

	dst := filepath.Join(out, ".github", "workflows")
	if err := os.MkdirAll(dst, 0o755); err != nil {
		panic(err)
	}

	rendered := map[string]bool{}
	for _, b := range bundles {
		for _, file := range b.Files {
			renderFile(dst, file, data)
			rendered[file] = true
			fmt.Fprintln(os.Stderr, "rendered", repo, "→", file)
		}
	}

	entries, err := os.ReadDir(dst)
	if err != nil {
		panic(err)
	}
	for _, e := range entries {
		if e.IsDir() || rendered[e.Name()] {
			continue
		}
		p := filepath.Join(dst, e.Name())
		b, err := os.ReadFile(p)
		if err != nil {
			panic(err)
		}
		if !strings.HasPrefix(string(b), marker) {
			continue
		}
		if err := os.Remove(p); err != nil {
			panic(err)
		}
		fmt.Fprintln(os.Stderr, "removed", repo, "→", e.Name())
	}
}

func renderFile(dst, file string, data map[string]any) {
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
	if _, err := fmt.Fprintf(f, header, file); err != nil {
		panic(err)
	}
	if err := tmpl.Execute(f, data); err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
}
