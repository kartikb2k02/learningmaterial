// seocheck audits the built site (public/) the way a crawler sees it.
//
//	hugo && go run ./tools/seocheck             # audit public/
//	go run ./tools/seocheck -external           # ...and request every outbound link
//	go run ./tools/seocheck -dir public -v      # list every finding, not just the first few
//
// It checks, for every HTML page: title (missing, duplicate, length), meta description
// (missing, duplicate, length), canonical (missing, duplicate, mismatch), accidental noindex,
// exactly one <h1>, heading order, breadcrumbs, JSON-LD that parses, image alt text, and
// internal links that resolve. Site-wide it checks robots.txt, that the sitemap and the pages
// agree, and that no page is an orphan (no inbound links) or nearly one.
//
// Exit status is 1 when any error is found; warnings do not fail the run.
package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"html"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	reTitle     = regexp.MustCompile(`(?s)<title>(.*?)</title>`)
	reDesc      = regexp.MustCompile(`<meta name="?description"? content="([^"]*)"`)
	reCanonical = regexp.MustCompile(`<link rel="?canonical"? href="?([^" >]+)`)
	reRobots    = regexp.MustCompile(`<meta name="?robots"? content="([^"]*)"`)
	reHeading   = regexp.MustCompile(`(?s)<h([1-6])[\s>]`)
	reLink      = regexp.MustCompile(`<a\s[^>]*?href="?([^" >]+)`)
	reID        = regexp.MustCompile(`\sid="?([^" >]+)`)
	reImg       = regexp.MustCompile(`<img\s[^>]*>`)
	reAlt       = regexp.MustCompile(`\salt=("[^"]+"|[^" >]+)`)
	reLD        = regexp.MustCompile(`(?s)<script type="?application/ld\+json"?>(.*?)</script>`)
	reMain      = regexp.MustCompile(`(?s)<main[\s>].*?</main>`)
)

type page struct {
	file, path           string // file on disk, URL path ("/a/b/")
	title, desc, canon   string
	noindex              bool
	ids                  map[string]bool
	links                []string // raw hrefs
	bodyLinks            int      // links inside <main>
	inbound, bodyInbound int
}

type finding struct{ sev, check, where, msg string }

func main() {
	dir := flag.String("dir", "public", "the built site")
	external := flag.Bool("external", false, "also request outbound links (slow, needs network)")
	verbose := flag.Bool("v", false, "list every finding")
	flag.Parse()

	var out []finding
	add := func(sev, check, where, msg string) { out = append(out, finding{sev, check, where, msg}) }

	// ── read every page ──────────────────────────────────────────────────────
	pages := map[string]*page{}
	var order []string
	err := filepath.WalkDir(*dir, func(f string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(f, ".html") {
			return err
		}
		rel, _ := filepath.Rel(*dir, f)
		p := "/" + filepath.ToSlash(strings.TrimSuffix(rel, "index.html"))
		pages[p] = &page{file: f, path: p}
		order = append(order, p)
		return nil
	})
	if err != nil || len(pages) == 0 {
		fmt.Fprintf(os.Stderr, "seocheck: no pages under %q (run `hugo` first): %v\n", *dir, err)
		os.Exit(2)
	}
	sort.Strings(order)

	base := ""
	titles, descs, canons := map[string][]string{}, map[string][]string{}, map[string][]string{}
	external404 := map[string][]string{}

	for _, path := range order {
		pg := pages[path]
		raw, _ := os.ReadFile(pg.file)
		src := string(raw)
		is404 := strings.HasSuffix(path, "404.html")

		if m := reRobots.FindStringSubmatch(src); m != nil && strings.Contains(m[1], "noindex") {
			pg.noindex = true
			if !is404 {
				add("error", "noindex", path, "page is marked noindex")
			}
		}
		if m := reTitle.FindStringSubmatch(src); m != nil {
			pg.title = html.UnescapeString(strings.TrimSpace(m[1]))
		}
		if m := reDesc.FindStringSubmatch(src); m != nil {
			pg.desc = html.UnescapeString(strings.TrimSpace(m[1]))
		}
		if m := reCanonical.FindStringSubmatch(src); m != nil {
			pg.canon = html.UnescapeString(m[1])
		}
		pg.ids = map[string]bool{}
		for _, m := range reID.FindAllStringSubmatch(src, -1) {
			if pg.ids[m[1]] {
				add("warn", "duplicate id", path, "#"+m[1])
			}
			pg.ids[m[1]] = true
		}
		for _, m := range reLink.FindAllStringSubmatch(src, -1) {
			pg.links = append(pg.links, html.UnescapeString(m[1]))
		}
		body := reMain.FindString(src)
		pg.bodyLinks = len(reLink.FindAllString(body, -1))
		if is404 {
			continue
		}

		// title
		switch n := len([]rune(pg.title)); {
		case n == 0:
			add("error", "title", path, "missing <title>")
		case n < 15:
			add("warn", "title", path, fmt.Sprintf("short title (%d chars): %q", n, pg.title))
		case n > 70:
			add("warn", "title", path, fmt.Sprintf("long title (%d chars): %q", n, pg.title))
		}
		titles[pg.title] = append(titles[pg.title], path)

		// description
		switch n := len([]rune(pg.desc)); {
		case n == 0:
			add("error", "description", path, "missing meta description")
		case n < 50:
			add("warn", "description", path, fmt.Sprintf("short description (%d chars): %q", n, pg.desc))
		case n > 170:
			add("warn", "description", path, fmt.Sprintf("long description (%d chars)", n))
		}
		descs[pg.desc] = append(descs[pg.desc], path)

		// canonical: present, absolute, https, self-referencing
		if pg.canon == "" {
			add("error", "canonical", path, "missing canonical")
		} else if u, err := url.Parse(pg.canon); err != nil || !u.IsAbs() {
			add("error", "canonical", path, "canonical is not an absolute URL: "+pg.canon)
		} else {
			if u.Scheme != "https" {
				add("warn", "canonical", path, "canonical is not https: "+pg.canon)
			}
			if u.Path != path {
				add("error", "canonical", path, "canonical points elsewhere: "+pg.canon)
			}
			if base == "" {
				base = u.Scheme + "://" + u.Host
			}
			canons[pg.canon] = append(canons[pg.canon], path)
		}

		// headings: one h1, no skipped levels
		last, h1 := 0, 0
		for _, m := range reHeading.FindAllStringSubmatch(body, -1) {
			l := int(m[1][0] - '0')
			if l == 1 {
				h1++
			}
			if last != 0 && l > last+1 {
				add("warn", "heading order", path, fmt.Sprintf("h%d follows h%d", l, last))
			}
			last = l
		}
		if h1 != 1 {
			add("error", "h1", path, fmt.Sprintf("%d <h1> elements in <main>", h1))
		}

		// breadcrumbs: visible nav + BreadcrumbList, on everything below the home page
		if path != "/" {
			if !strings.Contains(src, `aria-label="Breadcrumb"`) && !strings.Contains(src, `aria-label=Breadcrumb`) {
				add("warn", "breadcrumbs", path, "no visible breadcrumb nav")
			}
			if !strings.Contains(src, `"BreadcrumbList"`) {
				add("error", "breadcrumbs", path, "no BreadcrumbList structured data")
			}
		}

		// structured data parses and names a type
		lds := reLD.FindAllStringSubmatch(src, -1)
		if len(lds) == 0 {
			add("error", "structured data", path, "no JSON-LD")
		}
		for _, m := range lds {
			var v map[string]any
			if err := json.Unmarshal([]byte(m[1]), &v); err != nil {
				add("error", "structured data", path, "JSON-LD does not parse: "+err.Error())
			} else if v["@context"] == nil {
				add("error", "structured data", path, "JSON-LD has no @context")
			}
		}

		// images
		for _, img := range reImg.FindAllString(src, -1) {
			if !reAlt.MatchString(img) {
				add("error", "image alt", path, "image without alt text: "+img)
			}
		}

		// social card
		for _, tag := range []string{"og:title", "og:description", "og:url", "og:image", "twitter:card"} {
			if !strings.Contains(src, tag) {
				add("error", "social", path, "missing "+tag)
			}
		}
	}

	dup := func(check string, m map[string][]string) {
		for k, ps := range m {
			if k != "" && len(ps) > 1 {
				add("error", "duplicate "+check, ps[0], fmt.Sprintf("%q is shared by %d pages: %s", k, len(ps), strings.Join(ps, ", ")))
			}
		}
	}
	dup("title", titles)
	dup("description", descs)
	dup("canonical", canons)

	// ── links ────────────────────────────────────────────────────────────────
	weak := regexp.MustCompile(`(?i)>\s*(click here|read more|learn more|here|this link)\s*</a>`)
	for _, path := range order {
		pg := pages[path]
		raw, _ := os.ReadFile(pg.file)
		if m := weak.FindString(string(raw)); m != "" {
			add("warn", "anchor text", path, "vague link text: "+strings.TrimSpace(m))
		}
		seen := map[string]bool{}
		for _, href := range pg.links {
			u, err := url.Parse(href)
			if err != nil {
				add("error", "malformed url", path, href)
				continue
			}
			if u.Scheme == "mailto" || u.Scheme == "tel" || u.Scheme == "javascript" {
				continue
			}
			if u.IsAbs() && (base == "" || u.Scheme+"://"+u.Host != base) {
				external404[href] = append(external404[href], path)
				continue
			}
			target := u.Path
			if target == "" {
				target = path
			} else if !strings.HasPrefix(target, "/") {
				target = resolve(path, target)
			}
			tp, ok := pages[target]
			if !ok {
				// a file (css, json, image) rather than a page?
				if _, err := os.Stat(filepath.Join(*dir, filepath.FromSlash(target))); err == nil && !strings.HasSuffix(target, "/") {
					continue
				}
				if _, ok := pages[target+"/"]; ok {
					add("warn", "trailing slash", path, "link without trailing slash (redirects): "+href)
				} else {
					add("error", "broken link", path, href)
				}
				continue
			}
			if u.Fragment != "" && !tp.ids[u.Fragment] {
				if f, err := url.PathUnescape(u.Fragment); err != nil || !tp.ids[f] {
					add("warn", "broken anchor", path, href)
				}
			}
			if target != path && !seen[target] {
				seen[target] = true
				tp.inbound++
			}
		}
	}
	// inbound links from page bodies only (navigation links every page to everything)
	for _, path := range order {
		raw, _ := os.ReadFile(pages[path].file)
		seen := map[string]bool{}
		for _, m := range reLink.FindAllStringSubmatch(reMain.FindString(string(raw)), -1) {
			u, err := url.Parse(html.UnescapeString(m[1]))
			if err != nil || u.IsAbs() || u.Path == "" {
				continue
			}
			if tp, ok := pages[u.Path]; ok && u.Path != path && !seen[u.Path] {
				seen[u.Path] = true
				tp.bodyInbound++
			}
		}
	}
	for _, path := range order {
		pg := pages[path]
		if path == "/" || strings.HasSuffix(path, "404.html") {
			continue
		}
		if pg.inbound == 0 {
			add("error", "orphan page", path, "no page links here")
		} else if pg.bodyInbound < 2 {
			add("warn", "few inbound links", path, fmt.Sprintf("%d in-content link(s) point here", pg.bodyInbound))
		}
	}

	// ── robots.txt and sitemap ───────────────────────────────────────────────
	robots, err := os.ReadFile(filepath.Join(*dir, "robots.txt"))
	switch {
	case err != nil:
		add("error", "robots.txt", "/robots.txt", "missing")
	case !strings.Contains(string(robots), "Sitemap:"):
		add("error", "robots.txt", "/robots.txt", "does not reference the sitemap")
	default:
		for _, line := range strings.Split(string(robots), "\n") {
			if v, ok := strings.CutPrefix(strings.TrimSpace(line), "Disallow:"); ok {
				v = strings.TrimSpace(v)
				if v == "/" {
					add("error", "robots.txt", "/robots.txt", "blocks the whole site")
				}
				for _, asset := range []string{"/css", "/js", "/fonts", "/vendor"} {
					if strings.HasPrefix(v, asset) {
						add("error", "robots.txt", "/robots.txt", "blocks assets needed to render: "+v)
					}
				}
			}
		}
	}
	inSitemap := map[string]bool{}
	if sm, err := os.ReadFile(filepath.Join(*dir, "sitemap.xml")); err != nil {
		add("error", "sitemap", "/sitemap.xml", "missing")
	} else {
		var set struct {
			URLs []struct {
				Loc string `xml:"loc"`
			} `xml:"url"`
		}
		if err := xml.Unmarshal(sm, &set); err != nil {
			add("error", "sitemap", "/sitemap.xml", "does not parse: "+err.Error())
		}
		for _, u := range set.URLs {
			pu, err := url.Parse(u.Loc)
			if err != nil || !pu.IsAbs() {
				add("error", "sitemap", "/sitemap.xml", "not an absolute URL: "+u.Loc)
				continue
			}
			if inSitemap[pu.Path] {
				add("error", "sitemap", pu.Path, "listed twice")
			}
			inSitemap[pu.Path] = true
			pg, ok := pages[pu.Path]
			switch {
			case !ok:
				add("error", "sitemap", pu.Path, "in the sitemap but not built")
			case pg.noindex:
				add("error", "sitemap", pu.Path, "noindex page listed in the sitemap")
			case pg.canon != u.Loc:
				add("error", "sitemap", pu.Path, "sitemap URL differs from the canonical "+pg.canon)
			}
		}
		for _, path := range order {
			if !inSitemap[path] && !pages[path].noindex {
				add("error", "sitemap", path, "indexable page missing from the sitemap")
			}
		}
	}

	// ── outbound links (optional) ────────────────────────────────────────────
	if *external {
		checkExternal(external404, add)
	}

	// ── report ───────────────────────────────────────────────────────────────
	indexable := 0
	for _, p := range pages {
		if !p.noindex {
			indexable++
		}
	}
	fmt.Printf("SEO audit of %s\n\n", *dir)
	fmt.Printf("  pages            %d\n  indexable        %d\n  noindex          %d\n  in sitemap       %d\n", len(pages), indexable, len(pages)-indexable, len(inSitemap))
	fmt.Printf("  unique titles    %d\n  unique descs     %d\n  outbound links   %d distinct%s\n\n", len(titles), len(descs), len(external404), map[bool]string{true: " (checked)", false: " (not checked; use -external)"}[*external])

	groups := map[string][]finding{}
	var keys []string
	errs := 0
	for _, f := range out {
		k := f.sev + "  " + f.check
		if groups[k] == nil {
			keys = append(keys, k)
		}
		groups[k] = append(groups[k], f)
		if f.sev == "error" {
			errs++
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		fs := groups[k]
		fmt.Printf("%s — %d\n", k, len(fs))
		for i, f := range fs {
			if !*verbose && i == 5 {
				fmt.Printf("    … %d more (run with -v)\n", len(fs)-5)
				break
			}
			fmt.Printf("    %s  %s\n", f.where, f.msg)
		}
	}
	if len(out) == 0 {
		fmt.Println("no findings")
	}
	fmt.Printf("\n%d error(s), %d warning(s)\n", errs, len(out)-errs)
	if errs > 0 {
		os.Exit(1)
	}
}

// resolve turns a relative href into a site path, relative to the page it is on.
func resolve(from, href string) string {
	u, err := url.Parse("http://x" + from)
	if err != nil {
		return href
	}
	r, err := u.Parse(href)
	if err != nil {
		return href
	}
	return r.Path
}

func checkExternal(links map[string][]string, add func(sev, check, where, msg string)) {
	client := &http.Client{Timeout: 15 * time.Second}
	type res struct {
		href string
		err  string
	}
	jobs, results := make(chan string), make(chan res)
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for href := range jobs {
				req, _ := http.NewRequest("GET", href, nil)
				req.Header.Set("User-Agent", "Mozilla/5.0 (seocheck link audit)")
				resp, err := client.Do(req)
				switch {
				case err != nil:
					results <- res{href, err.Error()}
				case resp.StatusCode >= 400 && resp.StatusCode != 403 && resp.StatusCode != 429:
					results <- res{href, resp.Status}
				default:
					results <- res{href, ""}
				}
				if resp != nil {
					resp.Body.Close()
				}
			}
		}()
	}
	go func() {
		for href := range links {
			jobs <- href
		}
		close(jobs)
		wg.Wait()
		close(results)
	}()
	for r := range results {
		if r.err != "" {
			add("warn", "broken outbound link", links[r.href][0], r.href+" — "+r.err)
		}
	}
}
