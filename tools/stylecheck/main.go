// stylecheck enforces the house rule that prose and bullets carry no dashes.
//
//	go run ./tools/stylecheck product-management metrics-and-models
//
// It reads every .md file under the given paths and reports a dash left in prose.
// Four things are exempt, because tooling or markdown requires them:
//
//	fenced code blocks        the programs, and the Level/Prerequisites metadata block
//	table delimiter rows      |---|---| is the only legal spelling
//	inline code spans         `01-foo.md`, `--flag`
//	link and image targets    the (…) half of [text](path), since filenames keep prefixes
//
// Everything else is prose, where a comma, colon, semicolon or bracket is wanted instead.
package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	fence     = regexp.MustCompile("^[ \t]*```")
	delimiter = regexp.MustCompile(`^[ \t]*\|[\s:|-]+\|[ \t]*$`)
	code      = regexp.MustCompile("`[^`]*`")
	target    = regexp.MustCompile(`\]\([^)]*\)`)
	anchor    = regexp.MustCompile(`^\s*\[[^\]]+\]:\s*\S+`)
)

func main() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}
	bad := 0
	for _, root := range roots {
		filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() || !strings.HasSuffix(path, ".md") {
				return nil
			}
			bad += check(path)
			return nil
		})
	}
	if bad > 0 {
		fmt.Fprintf(os.Stderr, "\nstylecheck: %d lines with a dash in prose\n", bad)
		os.Exit(1)
	}
	fmt.Println("stylecheck: clean")
}

func check(path string) int {
	f, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer f.Close()
	var (
		in  bool
		n   int
		bad int
	)
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	for sc.Scan() {
		n++
		line := sc.Text()
		if fence.MatchString(line) {
			in = !in
			continue
		}
		if in || delimiter.MatchString(line) || anchor.MatchString(line) {
			continue
		}
		stripped := target.ReplaceAllString(code.ReplaceAllString(line, ""), "]()")
		if strings.Contains(stripped, "-") {
			bad++
			fmt.Printf("%s:%d: %s\n", path, n, strings.TrimSpace(line))
		}
	}
	return bad
}
