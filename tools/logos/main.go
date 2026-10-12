// logos vendors the small icons that diagrams use into static/logos/, so the site stays offline.
//
//	go run ./tools/logos            # fetch everything listed below (needs network, run rarely)
//
// A diagram refers to an icon by name between colons — ":kubernetes: Kubernetes" — and
// assets/js/main.js draws static/logos/<name>.svg in its place.
//
// Three kinds of icon end up in the folder:
//
//	brand     a product's mark from Simple Icons (CC0), tinted with the brand colour
//	monogram  a lettered tile, for projects Simple Icons does not carry
//	glyph     a generic pictogram from Lucide (ISC), named "i-<name>"
//
// Brand marks are trademarks of their owners and are used here only to identify the product.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	simpleIcons = "https://cdn.jsdelivr.net/npm/simple-icons@16.34.0"
	lucide      = "https://cdn.jsdelivr.net/npm/lucide-static@1.51.0/icons"
	out         = "static/logos"
	ink         = "#1e1e1e"
)

// Simple Icons slugs.
var brands = strings.Fields(`
	kubernetes nvidia prometheus grafana opentelemetry envoyproxy istio redis postgresql apachekafka
	natsdotio temporal docker containerd helm argo terraform vault keycloak trivy falco cilium
	huggingface anthropic claude googlegemini meta mistralai ollama langchain langgraph mlflow ray
	qdrant milvus elasticsearch opensearch clickhouse minio cloudflare nginx kong github gitlab
	python go pytorch onnx modelcontextprotocol okta auth0 snyk linux ceph duckdb jaeger datadog
	sentry pagerduty webassembly weightsandbiases harbor googlecloud pydantic amd intel google owasp
	cncf linuxfoundation knative etcd spdx hashicorp thanos githubcopilot cursor modal qemu
	googlechrome stripe notion jira gmail googledrive victoriametrics fluentbit deepseek crewai
	neo4j snowflake databricks apachespark apacheairflow backstage confluence dvc
	hubspot mixpanel posthog shopify zapier intercom zendesk linear figma webflow mailchimp
	googleanalytics looker metabase airtable asana trello miro razorpay shopware woocommerce
	ebay airbnb uber doordash`)

// Lettered tiles: name, letters, colour.
var monograms = [][3]string{
	{"vllm", "vL", "#30a2ff"}, {"sglang", "SG", "#7048e8"}, {"llm-d", "ld", "#0b7285"},
	{"kserve", "KS", "#326ce5"}, {"kueue", "Kq", "#326ce5"}, {"keda", "KE", "#3d6dd6"},
	{"litellm", "LL", "#0ca678"}, {"langfuse", "Lf", "#e8590c"}, {"opa", "OPA", "#566366"},
	{"spiffe", "SP", "#2f9e44"}, {"sigstore", "Sg", "#2a1e71"}, {"openai", "OA", "#1e1e1e"},
	{"gvisor", "gV", "#1971c2"}, {"firecracker", "FC", "#e8590c"}, {"kata", "Ka", "#f08c00"},
	{"e2b", "E2B", "#e67700"}, {"a2a", "A2A", "#1a73e8"}, {"agentgateway", "AG", "#7c3aed"},
	{"aws", "AWS", "#232f3e"}, {"azure", "Az", "#0078d4"}, {"mitre", "ATL", "#c92a2a"},
	{"nist", "NIST", "#343a40"}, {"slsa", "SL", "#2f9e44"}, {"cyclonedx", "CDX", "#a61e4d"},
	{"openssf", "SSF", "#12b76a"}, {"garak", "Gk", "#5f3dc4"}, {"pyrit", "PY", "#0078d4"},
	{"kyverno", "Ky", "#e8590c"}, {"semgrep", "Sm", "#12b76a"}, {"slack", "Sl", "#4a154b"},
	{"playwright", "Pw", "#2f9e44"}, {"tgi", "TGI", "#e67700"}, {"dynamo", "Dy", "#76b900"},
	{"slurm", "Sl", "#1971c2"}, {"lancedb", "La", "#e8590c"}, {"oauth", "OA2", "#343a40"},
	{"s3", "S3", "#2f9e44"}, {"otel-genai", "Gen", "#4f62ad"}, {"coco", "CoCo", "#c2255c"},
	{"llamaguard", "LG", "#0668e1"}, {"presidio", "Pr", "#0078d4"}, {"tetragon", "Te", "#e67700"},
	{"gong", "Gg", "#6c2bd9"}, {"outreach", "Or", "#5951ff"}, {"clay", "Cy", "#1e1e1e"},
	{"pendo", "Pn", "#ec2d8a"}, {"chargebee", "Cb", "#ff3300"}, {"productboard", "Pb", "#0d7ff2"},
	{"marketo", "Mk", "#5c4c9f"}, {"sixsense", "6s", "#ff4a00"}, {"clearbit", "Cl", "#4f46e5"},
	{"heap", "Hp", "#1a6fdb"}, {"statsig", "St", "#194b7d"}, {"optimizely", "Op", "#0037ff"},
	{"maxio", "Mx", "#0f766e"}, {"avalara", "Av", "#ff5100"}, {"vanta", "Va", "#6456f0"},
	{"drata", "Dr", "#5c4df5"}, {"ramp", "Rp", "#1e1e1e"}, {"brex", "Bx", "#f56b4b"},
	{"salesforce", "SF", "#00a1e0"}, {"amplitude", "Am", "#1f8ded"}, {"segment", "Sg", "#52bd94"},
	{"braze", "Bz", "#ff6b4a"}, {"klaviyo", "Kv", "#1f2d3d"}, {"dbt", "dbt", "#ff694b"},
	{"tableau", "Tb", "#e97627"}, {"canva", "Cv", "#00c4cc"}, {"twilio", "Tw", "#f22f46"},
	{"plaid", "Pl", "#1e1e1e"},
}

// Lucide icon names, saved as i-<name>.svg.
var glyphs = strings.Fields(`
	user users bot wrench database cpu shield-check shield-alert lock key-round globe file-text brain
	layers network server box scale eye bug skull siren workflow search hard-drive cloud terminal
	message-square zap gauge coins scroll-text fingerprint radar activity package git-branch clock
	mail triangle-alert circuit-board memory-stick router funnel list-checks hand flask-conical
	book-open code monitor smartphone building-2 ban check route split merge recycle archive
	trending-up trending-down target handshake megaphone shopping-cart credit-card receipt
	landmark store ticket crown filter calendar percent repeat rocket flag lightbulb puzzle
	gavel stethoscope factory truck pen-tool layout-dashboard thumbs-up thumbs-down banknote
	piggy-bank wallet badge-check tag inbox phone presentation sliders-horizontal compass
	door-open hourglass heart-crack user-plus user-minus user-check circle-help circle-slash
	arrow-up-right arrow-down-right arrow-left-right notebook-pen stamp shield scissors`)

var rePath = regexp.MustCompile(`<path d="([^"]+)"`)

func main() {
	check(os.MkdirAll(out, 0o755))

	var index []struct{ Title, Slug, Hex string }
	check(json.Unmarshal(get(simpleIcons+"/data/simple-icons.json"), &index))
	hex := map[string]string{}
	for _, e := range index {
		if e.Slug != "" {
			hex[e.Slug] = e.Hex
		}
	}
	for _, slug := range brands {
		m := rePath.FindSubmatch(get(simpleIcons + "/icons/" + slug + ".svg"))
		if m == nil {
			fail("no path in %s", slug)
		}
		colour := "#" + hex[slug]
		if hex[slug] == "" || luminance(hex[slug]) > 0.72 || luminance(hex[slug]) < 0.03 {
			colour = ink // a white or black mark is drawn in the diagram ink
		}
		write(slug, fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path fill="%s" d="%s"/></svg>`, colour, m[1]))
	}
	for _, m := range monograms {
		size := map[int]string{1: "13", 2: "11.5", 3: "8.6", 4: "6.8"}[len(m[1])]
		write(m[0], fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><rect x="1" y="1" width="22" height="22" rx="5" fill="%s"/><text x="12" y="12" dy=".36em" text-anchor="middle" font-family="Helvetica,Arial,sans-serif" font-weight="700" font-size="%s" fill="#fff">%s</text></svg>`, m[2], size, m[1]))
	}
	for _, g := range glyphs {
		svg := string(get(lucide + "/" + g + ".svg"))
		svg = regexp.MustCompile(`(?s)<!--.*?-->\s*`).ReplaceAllString(svg, "")
		svg = regexp.MustCompile(`\s+class="[^"]*"|\s+width="24"|\s+height="24"`).ReplaceAllString(svg, "")
		svg = strings.ReplaceAll(svg, `stroke="currentColor"`, `stroke="`+ink+`"`)
		write("i-"+g, strings.Join(strings.Fields(svg), " "))
	}
	fmt.Printf("%d brand marks, %d monograms, %d glyphs in %s\n", len(brands), len(monograms), len(glyphs), out)
}

func luminance(hex string) float64 {
	n, _ := strconv.ParseUint(hex, 16, 32)
	return (0.2126*float64(n>>16&255) + 0.7152*float64(n>>8&255) + 0.0722*float64(n&255)) / 255
}

func get(url string) []byte {
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fail("%s: %s", url, resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	check(err)
	return b
}

func write(name, svg string) {
	check(os.WriteFile(filepath.Join(out, name+".svg"), []byte(svg+"\n"), 0o644))
}

func check(err error) {
	if err != nil {
		fail("%v", err)
	}
}

func fail(format string, a ...any) {
	fmt.Fprintf(os.Stderr, "logos: "+format+"\n", a...)
	os.Exit(1)
}
