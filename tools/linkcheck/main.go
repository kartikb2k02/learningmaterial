// linkcheck resolves every relative markdown link, because lessons must read on GitHub as
// well as on the site.
//
//	go run ./tools/linkcheck                 # report dangling links
//	go run ./tools/linkcheck -pending        # list them grouped by the topic that owns them
//
// Absolute URLs, anchors and mailto links are ignored. A link into a topic folder that has
// not been written yet is still a dangling link; the -pending view is there so forward
// references can be tracked rather than forgotten.
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var (
	pending = flag.Bool("pending", false, "group dangling links by the folder that would satisfy them")
	link    = regexp.MustCompile(`\[[^\]]*\]\(([^)\s]+)(?:\s+"[^"]*")?\)`)
	fence   = regexp.MustCompile("(?s)```.*?```")
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	type miss struct{ from, to string }
	var misses []miss
	for _, root := range roots {
		filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() || !strings.HasSuffix(path, ".md") || strings.HasPrefix(path, "public/") {
				return nil
			}
			src, err := os.ReadFile(path)
			if err != nil {
				return nil
			}
			body := fence.ReplaceAllString(string(src), "")
			for _, m := range link.FindAllStringSubmatch(body, -1) {
				t := m[1]
				if strings.HasPrefix(t, "#") || strings.Contains(t, "://") || strings.HasPrefix(t, "mailto:") {
					continue
				}
				t = strings.SplitN(t, "#", 2)[0]
				if t == "" {
					continue
				}
				if _, err := os.Stat(filepath.Join(filepath.Dir(path), t)); err != nil {
					misses = append(misses, miss{path, t})
				}
			}
			return nil
		})
	}
	if len(misses) == 0 {
		fmt.Println("linkcheck: clean")
		return
	}
	if *pending {
		byDir := map[string][]string{}
		for _, m := range misses {
			target := filepath.Clean(filepath.Join(filepath.Dir(m.from), m.to))
			byDir[filepath.Dir(target)] = append(byDir[filepath.Dir(target)], filepath.Base(target))
		}
		dirs := make([]string, 0, len(byDir))
		for d := range byDir {
			dirs = append(dirs, d)
		}
		sort.Strings(dirs)
		for _, d := range dirs {
			files := byDir[d]
			sort.Strings(files)
			fmt.Printf("%s\n", d)
			seen := map[string]bool{}
			for _, f := range files {
				if !seen[f] {
					seen[f] = true
					fmt.Printf("    %s\n", f)
				}
			}
		}
		fmt.Printf("\n%d dangling links across %d folders not yet written\n", len(misses), len(byDir))
		return
	}
	for _, m := range misses {
		fmt.Printf("%s -> %s\n", m.from, m.to)
	}
	fmt.Fprintf(os.Stderr, "\nlinkcheck: %d dangling links\n", len(misses))
	os.Exit(1)
}
