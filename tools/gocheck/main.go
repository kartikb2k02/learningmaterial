// gocheck keeps the Go in the lessons honest.
//
// It finds every ```go block that is a complete program (contains "package main"),
// formats it with gofmt (rewriting the lesson in place), and runs `go vet` on it.
// With -run it also executes each program that needs no arguments, network or GPU.
//
//	go run ./tools/gocheck inference-engineering gpu-engineering
//	go run ./tools/gocheck -run inference-engineering/I-fundamentals
package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"go/format"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	run     = flag.Bool("run", false, "also run programs that need no arguments, network or GPU")
	fence   = regexp.MustCompile("(?m)^([ \t]*)```go[ \t]*\n((?s:.*?))\n[ \t]*```[ \t]*$")
	needEnv = regexp.MustCompile(`os\.Args|net/http|"net"|import "C"|os\.Stdin|go-nvml`)
)

func main() {
	flag.Parse()
	tmp, err := os.MkdirTemp("", "gocheck")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmp)

	total, bad := 0, 0
	for _, root := range flag.Args() {
		filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() || !strings.HasSuffix(path, ".md") {
				return nil
			}
			src, _ := os.ReadFile(path)
			n := 0
			out := fence.ReplaceAllFunc(src, func(m []byte) []byte {
				sub := fence.FindSubmatch(m)
				indent, code := string(sub[1]), sub[2]
				if indent != "" || !bytes.Contains(code, []byte("package main")) {
					return m
				}
				n++
				total++
				name := fmt.Sprintf("%s #%d", path, n)
				pretty, err := format.Source(code)
				if err != nil {
					bad++
					fmt.Printf("FAIL %s\n  gofmt: %v\n", name, err)
					return m
				}
				if msg := check(tmp, pretty); msg != "" {
					bad++
					fmt.Printf("FAIL %s\n%s\n", name, msg)
				}
				return []byte("```go\n" + strings.TrimRight(string(pretty), "\n") + "\n```")
			})
			if !bytes.Equal(out, src) {
				os.WriteFile(path, out, 0o644)
			}
			return nil
		})
	}
	fmt.Printf("%d programs checked, %d failed\n", total, bad)
	if bad > 0 {
		os.Exit(1)
	}
}

// check vets (and optionally runs) one program; it returns "" when everything is fine.
func check(tmp string, code []byte) string {
	dir, _ := os.MkdirTemp(tmp, "p")
	os.WriteFile(filepath.Join(dir, "main.go"), code, 0o644)
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module lesson\n\ngo 1.22\n"), 0o644)
	if bytes.Contains(code, []byte(`import "C"`)) || bytes.Contains(code, []byte("github.com/")) {
		return "" // needs a C toolchain / CUDA / third-party module: formatted only
	}
	if out, err := cmd(dir, 2*time.Minute, "go", "vet", "."); err != nil {
		return indent(out)
	}
	if *run && !needEnv.Match(code) {
		if out, err := cmd(dir, time.Minute, "go", "run", "."); err != nil {
			return indent(out + err.Error())
		}
	}
	return ""
}

func cmd(dir string, limit time.Duration, name string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), limit)
	defer cancel()
	c := exec.CommandContext(ctx, name, args...)
	c.Dir = dir
	out, err := c.CombinedOutput()
	return string(out), err
}

func indent(s string) string {
	return "  " + strings.ReplaceAll(strings.TrimSpace(s), "\n", "\n  ")
}
