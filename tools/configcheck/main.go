// configcheck catches the two ways config.yaml goes wrong silently.
//
//	go run ./tools/configcheck
//
// Neither failure produces an error. Hugo parses the file, the site builds, and the
// text is quietly truncated, so the only symptom is a blurb or a summary that stops
// early on the rendered page.
//
//	flow mappings   inside { a: 1, b: 2 } a comma starts a new key, so
//	                "blurb: Start here, and learn the job." keeps only "Start here"
//	unquoted colons "summary: Building AI products: context and agents" is a nested
//	                mapping, not a string, and fails the build with a cryptic message
//
// The rule is therefore: block mappings only, and quote any scalar containing a colon.
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	scalar = regexp.MustCompile(`^(\s*)(?:-\s+)?([A-Za-z_][\w-]*):\s+(\S.*)$`)
	colon  = regexp.MustCompile(`:\s`)
)

func main() {
	path := "config.yaml"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "configcheck: %v\n", err)
		os.Exit(2)
	}
	defer f.Close()

	bad, n := 0, 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		n++
		line := sc.Text()
		if i := strings.Index(line, "#"); i == 0 || (i > 0 && strings.TrimSpace(line[:i]) == "") {
			continue
		}
		if strings.Contains(line, "{") {
			bad++
			fmt.Printf("%s:%d: flow mapping. A comma inside { } starts a new key and truncates the value.\n  %s\n", path, n, strings.TrimSpace(line))
			continue
		}
		m := scalar.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		v := strings.TrimSpace(m[3])
		if strings.HasPrefix(v, `"`) || strings.HasPrefix(v, "'") || strings.HasPrefix(v, "[") || strings.HasPrefix(v, "#") {
			continue
		}
		if i := strings.Index(v, " #"); i > 0 {
			v = strings.TrimSpace(v[:i])
		}
		if colon.MatchString(v) {
			bad++
			fmt.Printf("%s:%d: unquoted colon in a scalar. Wrap the value in double quotes.\n  %s\n", path, n, strings.TrimSpace(line))
		}
	}
	if bad > 0 {
		fmt.Fprintf(os.Stderr, "\nconfigcheck: %d problems\n", bad)
		os.Exit(1)
	}
	fmt.Println("configcheck: clean")
}
