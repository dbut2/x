package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic("Usage: finder <target file> [skip list]")
	}

	target := os.Args[1]

	var skips []string
	if len(os.Args) > 2 {
		for _, a := range os.Args[2:] {
			skips = append(skips, strings.Split(a, ",")...)
		}
	}

	dirs := slices.DeleteFunc(findDirs(target), func(s string) bool {
		return slices.Contains(skips, s)
	})

	output := map[string][]string{"dir": dirs}

	b, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func findDirs(target string) []string {
	files := find(target)
	dirs := make([]string, 0, len(files))
	for _, file := range files {
		dirs = append(dirs, filepath.Dir(strings.TrimSuffix(file, target)))
	}
	slices.Sort(dirs)
	return dirs
}

func find(target string) []string {
	files := []string{}
	_ = filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if path == target || strings.HasSuffix(path, "/"+target) {
			files = append(files, path)
		}
		return nil
	})
	slices.Sort(files)
	return files
}
