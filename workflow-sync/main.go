package main

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
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
	Gomods  []string `yaml:"gomods"`
	Remove  []string `yaml:"remove"`
	RunsOn  string   `yaml:"runs-on"`
	Retries int      `yaml:"retries"`
	Branch  string   `yaml:"branch"`
}

type bundle struct {
	Files  []string `yaml:"files"`
	Checks []string `yaml:"checks"`
}

type manifest struct {
	Bundles map[string]bundle `yaml:"bundles"`
	Always  []string          `yaml:"always"`
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
		type entry struct {
			Repo   string `json:"repo"`
			Branch string `json:"branch,omitempty"`
		}
		include := make([]entry, 0, len(repos))
		for _, repo := range repos {
			include = append(include, entry{Repo: repo, Branch: m.Targets[repo].Branch})
		}
		b, err := json.Marshal(map[string][]entry{"include": include})
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

	for _, name := range t.Bundles {
		if slices.Contains(m.Always, name) {
			panic("target lists an always bundle: " + repo + " " + name)
		}
	}
	names := append(slices.Clone(m.Always), t.Bundles...)
	bundles := make([]bundle, 0, len(names))
	for _, name := range names {
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
	retries := t.Retries
	if retries == 0 {
		retries = 10
	}
	if t.Gomod != "" && len(t.Gomods) > 0 {
		panic("target sets both gomod and gomods: " + repo)
	}
	gomod := t.Gomod
	if len(t.Gomods) == 1 {
		gomod = t.Gomods[0]
	}
	if gomod == "" {
		gomod = "go.mod"
	}
	gomodDir := path.Dir(gomod)
	modules := []string{gomodDir}
	var dirs []string
	if len(t.Gomods) > 1 {
		for _, g := range t.Gomods {
			dirs = append(dirs, path.Dir(g))
		}
		modules = dirs
		gomod = "${{ matrix.dir }}/go.mod"
		gomodDir = "${{ matrix.dir }}"
	}
	data := map[string]any{"Checks": checks, "Gomod": gomod, "GomodDir": gomodDir, "Dirs": dirs, "Modules": modules, "RunsOn": runsOn, "Retries": retries}

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

	for _, file := range t.Remove {
		if rendered[file] {
			panic("remove lists a rendered file: " + file)
		}
		err := os.Remove(filepath.Join(dst, file))
		if errors.Is(err, fs.ErrNotExist) {
			continue
		}
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(os.Stderr, "removed", repo, "→", file)
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
