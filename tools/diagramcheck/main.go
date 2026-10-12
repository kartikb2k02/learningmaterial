// diagramcheck validates the Mermaid blocks in lessons before a browser ever sees them.
//
//	go run ./tools/diagramcheck product-management metrics-and-models
//
// It checks the things that silently render wrong rather than loudly failing:
//
//	logo tokens      ":stripe:" only becomes a mark if static/logos/stripe.svg exists,
//	                 otherwise the page prints the colons literally
//	class names      only compute, memory, queue, io, warn and neutral are styled by
//	                 assets/js/main.js; any other name renders unstyled
//	HTML in labels   a label containing < must be quoted, and its <b> and <small> tags
//	                 must be balanced, or Mermaid drops or mangles the label
//	reserved ids     a node called "end" breaks the parser
//	label quoting    a label that opens with a quote must close it before the
//	                 bracket, or Mermaid swallows the rest of the line
//	classed nodes    every id in a class statement must be a node declared in the diagram
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	block    = regexp.MustCompile("(?s)```mermaid\n(.*?)\n```")
	logo     = regexp.MustCompile(`:([a-z0-9][a-z0-9-]*):`)
	classSt  = regexp.MustCompile(`(?m)^\s*class\s+([A-Za-z0-9_,\s]+?)\s+([a-z]+)\s*$`)
	label    = regexp.MustCompile(`(?:\[\(|\[|\(|\{)"?(.*?)"?(?:\)\]|\]|\)|\})`)
	nodeDecl = regexp.MustCompile(`(?:^|[\s|>-])([A-Za-z][A-Za-z0-9_]*)\s*(?:\[\(|\[|\(\(|\(|\{)`)
	htmlTag  = regexp.MustCompile(`</?([a-z]+)\s*/?>`)
	roles    = map[string]bool{"compute": true, "memory": true, "queue": true, "io": true, "warn": true, "neutral": true}
	headers  = []string{"flowchart", "graph", "sequenceDiagram", "stateDiagram", "erDiagram", "journey", "gantt", "pie", "mindmap", "timeline", "quadrantChart", "block"}
)

func main() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}
	logos := vendored()
	bad := 0
	for _, root := range roots {
		filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() || !strings.HasSuffix(path, ".md") {
				return nil
			}
			bad += check(path, logos)
			return nil
		})
	}
	if bad > 0 {
		fmt.Fprintf(os.Stderr, "\ndiagramcheck: %d problems\n", bad)
		os.Exit(1)
	}
	fmt.Println("diagramcheck: clean")
}

func vendored() map[string]bool {
	m := map[string]bool{}
	entries, err := os.ReadDir("static/logos")
	if err != nil {
		fmt.Fprintln(os.Stderr, "diagramcheck: cannot read static/logos; run from the repository root")
		os.Exit(2)
	}
	for _, e := range entries {
		m[strings.TrimSuffix(e.Name(), ".svg")] = true
	}
	return m
}

func check(path string, logos map[string]bool) int {
	src, err := os.ReadFile(path)
	if err != nil {
		return 0
	}
	bad := 0
	for i, m := range block.FindAllStringSubmatch(string(src), -1) {
		for _, p := range problems(m[1], logos) {
			bad++
			fmt.Printf("%s: diagram %d: %s\n", path, i+1, p)
		}
	}
	return bad
}

func problems(d string, logos map[string]bool) []string {
	var out []string
	lines := strings.Split(strings.TrimSpace(d), "\n")

	known := false
	for _, h := range headers {
		if strings.HasPrefix(strings.TrimSpace(lines[0]), h) {
			known = true
		}
	}
	if !known {
		out = append(out, fmt.Sprintf("unknown diagram type %q", strings.TrimSpace(lines[0])))
	}

	for _, t := range logo.FindAllStringSubmatch(d, -1) {
		if !logos[t[1]] {
			out = append(out, fmt.Sprintf("logo %q has no file in static/logos", t[1]))
		}
	}

	declared := map[string]bool{}
	for _, n := range nodeDecl.FindAllStringSubmatch(d, -1) {
		declared[n[1]] = true
	}
	if declared["end"] {
		out = append(out, `node id "end" breaks the Mermaid parser`)
	}
	for _, s := range regexp.MustCompile(`(?m)^\s*subgraph\s+([A-Za-z][A-Za-z0-9_]*)`).FindAllStringSubmatch(d, -1) {
		declared[s[1]] = true
	}

	for _, l := range strings.Split(d, "\n") {
		for _, open := range []string{`["`, `("`, `{"`, `[("`} {
			i := strings.Index(l, open)
			if i < 0 {
				continue
			}
			rest := l[i+len(open):]
			j := strings.Index(rest, `"`)
			k := strings.IndexAny(rest, "]})")
			if j < 0 || (k >= 0 && k < j) {
				out = append(out, fmt.Sprintf("label opens with a quote and never closes it: %s", trim(l)))
			}
		}
	}

	for _, c := range classSt.FindAllStringSubmatch(d, -1) {
		if !roles[c[2]] {
			out = append(out, fmt.Sprintf("class %q is not one of the six styled roles", c[2]))
		}
		for _, id := range strings.Split(c[1], ",") {
			if id = strings.TrimSpace(id); id != "" && !declared[id] {
				out = append(out, fmt.Sprintf("class statement names %q, which is not a node in this diagram", id))
			}
		}
	}

	for _, l := range label.FindAllStringSubmatch(d, -1) {
		text := l[1]
		if !strings.Contains(text, "<") {
			continue
		}
		raw := l[0]
		if !strings.Contains(raw, `"`) {
			out = append(out, fmt.Sprintf("label contains HTML but is not quoted: %s", trim(raw)))
		}
		depth := map[string]int{}
		for _, t := range htmlTag.FindAllStringSubmatch(text, -1) {
			if strings.HasSuffix(t[0], "/>") || t[1] == "br" {
				continue
			}
			if strings.HasPrefix(t[0], "</") {
				depth[t[1]]--
			} else {
				depth[t[1]]++
			}
		}
		for tag, n := range depth {
			if n != 0 {
				out = append(out, fmt.Sprintf("unbalanced <%s> in label: %s", tag, trim(text)))
			}
		}
	}
	return out
}

func trim(s string) string {
	s = strings.Join(strings.Fields(s), " ")
	if len(s) > 72 {
		s = s[:72] + "…"
	}
	return s
}
