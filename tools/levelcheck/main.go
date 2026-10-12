// levelcheck verifies that every lesson's Level line matches the level its topic
// is listed under in config.yaml.
//
//	go run ./tools/levelcheck
//
// The site reads two different things and shows both as chips on a lesson page: the
// track, which comes from the levels lists in config.yaml, and the difficulty, which
// comes from the "Level:" line in the lesson's own metadata block. When a topic moves
// between levels the config changes and the lessons do not, so the page prints two
// contradictory chips and nothing errors.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	levelsBlock = regexp.MustCompile(`(?m)^      levels:\n((?:        \w+:.*\n)+)`)
	levelLine   = regexp.MustCompile(`(?m)^        (\w+):\s*\[(.*)\]`)
	lessonLevel = regexp.MustCompile(`(?m)^Level:[ \t]*(.+?)[ \t]*$`)
	lessonFile  = regexp.MustCompile(`^\d\d-.*\.md$`)
)

func main() {
	course := "product-management"
	if len(os.Args) > 1 {
		course = os.Args[1]
	}
	cfg, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "levelcheck: %v\n", err)
		os.Exit(2)
	}
	m := levelsBlock.FindSubmatch(cfg)
	if m == nil {
		fmt.Fprintln(os.Stderr, "levelcheck: no levels block found in config.yaml")
		os.Exit(2)
	}
	want := map[string]string{}
	for _, l := range levelLine.FindAllStringSubmatch(string(m[1]), -1) {
		name := strings.Title(l[1])
		for _, t := range strings.Split(l[2], ",") {
			if t = strings.TrimSpace(t); t != "" {
				want[t] = name
			}
		}
	}

	bad := 0
	entries, _ := os.ReadDir(course)
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		expect, listed := want[e.Name()]
		files, _ := os.ReadDir(filepath.Join(course, e.Name()))
		for _, f := range files {
			if !lessonFile.MatchString(f.Name()) {
				continue
			}
			path := filepath.Join(course, e.Name(), f.Name())
			body, err := os.ReadFile(path)
			if err != nil {
				continue
			}
			got := lessonLevel.FindSubmatch(body)
			switch {
			case !listed:
				// A folder outside every level is a Build shelf item and has no track.
				if got != nil {
					bad++
					fmt.Printf("%s: has a Level line but its folder is not listed under any level\n", path)
				}
			case got == nil:
				bad++
				fmt.Printf("%s: no Level line, expected %q\n", path, expect)
			case string(got[1]) != expect:
				bad++
				fmt.Printf("%s: Level is %q, config puts this topic under %q\n", path, got[1], expect)
			}
		}
	}
	if bad > 0 {
		fmt.Fprintf(os.Stderr, "\nlevelcheck: %d lessons disagree with config.yaml\n", bad)
		os.Exit(1)
	}
	fmt.Println("levelcheck: clean")
}
