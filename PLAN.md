# Plan of Action: AI Product

**Site:** product.kartikbhalerao.in · built on the Hugo shell from `riteshsonawane1372/Learning-materials`

The subject is **how modern AI products are actually built.** The full topic and lesson map is in
[`COURSE-MAP.md`](COURSE-MAP.md); this file holds the mechanics: how the shell works, the lesson
template, the checks, and the build order.

---

## 0. What I found in the reference repo (this changes the plan)

It is a **Hugo static site** where *content is plain markdown folders at the repo root* and
the site is generated from them by a content adapter (`content/_content.gotmpl`).

**The good news: you genuinely do not need to touch the codebase.** To add a course you:

1. Create a folder at the repo root with markdown in it.
2. Add ~10 lines to `config.yaml` (`params.courses` + one `module.mounts` line).
3. Optionally add SEO entries to `data/seo.yaml`.

`layouts/`, `assets/`, `static/fonts`, `tools/` are never edited. That's the whole integration.

**Two corrections to my first assumptions:**

* **Diagrams are Mermaid, not `.excalidraw` files.** The repo vendors `mermaid.min.js` and the
  **Excalifont** typeface, and `layouts/_markup/render-codeblock-mermaid.html` + `assets/js/main.js`
  render ```` ```mermaid ```` blocks in a hand sketched, Excalidraw *like* style with an **Expand**
  button. So diagrams are written as text inside the lesson file. This is strictly better for you:
  no binary assets, no export step, diffable in git, and they render on GitHub too.
* **There are no "modules".** The hierarchy is **course → topic → lesson**, and topics are assigned
  to one of five learning levels in `config.yaml`. The repo README says this explicitly:
  *"Nothing is numbered as a 'module'."*

**Conventions the site infers from the files themselves** (so lessons carry almost no frontmatter):

| Convention | Effect |
| --- | --- |
| `README.md` in a folder | that folder's landing page |
| folder inside a course | a **topic**; its order comes from the `levels:` list in `config.yaml` |
| `I-`, `II-` … folder prefix | topic ordering (Roman numerals, stripped from the title) |
| `01-`, `02-` … file prefix | lesson ordering |
| first `# Heading` in the file | the page title (and it's removed from the body) |
| ```` ```text ```` block with `Level: / Prerequisites: / Estimated difficulty: n/5 \| Estimated time: n min` | difficulty badge, prereqs, reading time |
| ```` ```mermaid ```` block | rendered hand drawn diagram |
| `:kubernetes:` / `:i-users:` inside a diagram label | draws `static/logos/<name>.svg` inline |
| `projects/` folder (not listed in `levels`) | appears under **"Build"** on the course page |
| first paragraph after the metadata block | the card blurb / meta description |

Links are relative `.md`, so **every file must read well both on GitHub and on the
site.** That's a real constraint on how you write.

---

## 1. Structure

Five courses. One per discipline, each with its own Foundations to Expert ladder, because the shell
puts every topic on a difficulty ladder and five business functions are not five difficulty levels.

| Course | Folder | Topics | Lessons |
|---|---|---|---|
| **AI Product Management** ⭐ | `ai-product-management` | 16 | ~107 |
| Metrics and Economics | `metrics-and-economics` | 8 | ~51 |
| AI Growth | `ai-growth` | 8 | ~47 |
| AI Go To Market | `ai-gtm` | 8 | ~46 |
| AI Operations | `ai-operations` | 8 | ~46 |

There is no separate AI course. The whole library is about AI products, so an `ai-engineering`
course inside it would be a tautology; that material is topics II to XII of the flagship.

Six rules keep it from duplicating itself.

1. **Metrics are defined once**, in `metrics-and-economics`. Everything else links. Without this
   you get four definitions of acquisition cost inside a year.
2. **Sector market shape is defined once**, in `ai-product-management/XIII-product-by-sector`,
   whose README carries the cross sector matrix. Other courses' sector topics carry function deltas
   only, 600 to 1,200 words each, and never restate buyer, cycle or regulation.
3. **Product shape is read before market.** `II-04 AI Product Shapes` (copilot, agent, autonomous
   workflow, AI native app, infrastructure, embedded) is a prerequisite for every sector lesson,
   because shape usually determines more than market. Skipping it is how sector content decays into
   industry commentary.
4. **Evaluation is owned by topic IX.** Other lessons reference the method rather than reteaching it.
5. **Cost arithmetic lives in `metrics-and-economics/VI-ai-cost-and-margin`.** Product lessons quote
   results and link.
6. Every course README ends with a **How it connects to the other courses** section.

Full lesson tables for all five courses: [`COURSE-MAP.md`](COURSE-MAP.md).

## 3. The lesson template, diagrams and the checks

All of this now lives in [`CONTENT_GUIDE.md`](CONTENT_GUIDE.md), which is the authoring contract:
the fixed page shape, the house style rule about dashes, the six Mermaid classes and the six
diagram grammars, the Go block rules, the four checks and the definition of done.

Four commands gate every lesson.

```bash
go run ./tools/gocheck -run ai-product-management   # format, vet and execute every Go block
go run ./tools/stylecheck ai-product-management     # a dash left in prose or a bullet
go run ./tools/diagramcheck ai-product-management   # logo tokens, class names, HTML in labels
go run ./tools/linkcheck -pending                   # dangling links, by unwritten folder
```

The one thing no tool catches: **the prose around a program must match what the program prints.**
Three of the first five programs written for this library had narratives their own output
contradicted, and running them was the only thing that found it.

## 4. Phase 0: fork and lock the look (one session, do this first)

1. Fork or clone to a new repo. Suggested names: `operator-material`, `gtm-material`,
   `operator-handbook`.
2. **Delete the seven AI infrastructure course folders.** Keep `layouts/ assets/ static/ content/ data/
   tools/ config.yaml go.mod .gitignore`.
3. Rewrite `config.yaml`: `baseURL`, `title`, `params.hero.title`, `params.footer`,
   `params.seo.*`, `params.author`, `params.github`, then `courses: []` and matching
   `module.mounts`. Also reword `params.levels` blurbs, the current ones are
   written for inference engineering. Suggested:

   ```yaml
   levels:
     - { key: foundations,  name: Foundations,  blurb: Build the mental model. }
     - { key: basic,        name: Basic,        blurb: Learn the core mechanisms. }
     - { key: intermediate, name: Intermediate, blurb: Run the playbooks. }
     - { key: advanced,     name: Advanced,     blurb: Operate at scale. }
     - { key: expert,       name: Expert,       blurb: Make the calls and defend them. }
   ```

4. Reduce `data/seo.yaml` to its header comment plus `pages: {}`.
5. Rewrite the root `README.md` as your own course table.
6. Check prerequisites: **Hugo extended, recent version** (the `_content.gotmpl` content adapter
   needs ≥ 0.126) and **Go** for `tools/`. Run `hugo server` and confirm an empty site builds.
7. Add the logos you need (step 3 above).
8. **Write exactly one golden lesson, completely**, plus its topic README and course README.
   Recommended: `metrics-and-models/IV-revenue-and-efficiency/01-cac-payback-and-ltv.md`, it
   exercises every element at once (a funnel diagram with logos, comparison tables, a real Go
   calculator, a real cost angle, a strong sector lens).
9. Verify on the rendered page: the diagram renders hand drawn and **Expand** works · level badge and
   reading time parsed from the ```` ```text ```` block · TOC built from your H2/H3 · breadcrumbs ·
   prev/next pager · relative `.md` links resolve · search finds it · it still reads fine on GitHub.

**Do not write lesson two until lesson one is finished.** Every later lesson is a copy of it; a
flaw in the golden lesson gets copied 200 times.

---

## 5. Build phases

| Phase | Deliverable | Lessons |
|---|---|---|
| 0 ✅ | Shell forked, courses registered, icon set vendored, four checks built, CI, topic I written | 8 |
| 0b | Folder renames, site retitle, prerequisite rewiring | 0 |
| 1 | Flagship **II, III, IV, V**: what a model changes, discovery, specs, primitives | ~26 |
| 2 | Flagship **VI, VII, VIII**: context, agents, UX | ~22 |
| 3 | Flagship **IX** evaluation, plus metrics **V** quality metrics | ~16 |
| 4 | Flagship **X, XI, XII**: pricing, guardrails, moats | ~20 |
| 5 | Metrics **I, II, III, IV, VI, VII, VIII** | ~43 |
| 6 | Flagship **XIII** remaining sectors, plus **XIV, XV, XVI** | ~24 |
| 7 | `ai-growth`, then `ai-gtm`, then `ai-operations` | ~139 |

About 297 lessons at full build, which is the same order of magnitude as the shell's own library and
the same multi month commitment. Phases 1 to 4 and 6 complete the flagship, which with topic I
already written is a publishable site on its own.

**Reach phase 3 quickly.** Evaluation is the topic the whole positioning rests on, and it is the
one subject where there is almost nothing good written for a product audience.

**Cadence:** six to eight lessons per working session, every program run and its output reconciled
against the prose before the lesson is committed. A topic ships complete or not at all.

## 6. Definition of done

In [`CONTENT_GUIDE.md`](CONTENT_GUIDE.md), as a checklist.

## 7. Risks

| Risk | Handling |
| --- | --- |
| Function × sector matrix explosion | Full treatment in `product-management/XIII` only; every other course carries short deltas |
| Sector lessons drifting into industry commentary | `IX-product-archetypes` is a prerequisite, shape before market |
| Generic MBA filler | Mandatory runnable numbers, a real tool landscape, and a worked decision per topic |
| Four definitions of CAC | `metrics-and-models` is the only definition site |
| AI bolted on as a chapter | `## The AI layer` in every lesson, *and* `ai-engineering` as a full course |
| 255 stubs in `main` | Definition of done; a topic ships complete or not at all |
| Benchmarks rotting | Date every number inline, like the reference repo's "checked on 3 October 2026" |
| `ai-engineering` drifting into infra | Hard scope boundary: it stops at the product surface. No serving, GPUs, or model training |
| Drifting into codebase edits | Only `config.yaml`, `data/seo.yaml`, course folders, and the `tools/logos` icon list |

---

## 8. Open decisions

1. **Site title.** "Operator" no longer describes the site. Proposed **AI Product**, hero line
   *How AI products are actually built*.
2. **Folder renames.** `product-management` to `ai-product-management`, and `metrics-and-models` to
   `metrics-and-economics`. Cheap now, expensive after a hundred lessons of cross links.
3. **Dropped material.** Classic platform and API product work, classic archetypes such as
   marketplace liquidity and network density, and non generative machine learning products lose
   their topics. Marketplace mechanics survive inside the sector lesson. If you want the rest, it
   belongs in a sixth `product-foundations` course rather than diluting the flagship.
4. **Should evaluation graduate to its own course?** It is eight lessons as topic IX. As a course it
   could carry twenty five, and it is the strongest claim the site has. The cost is fragmenting the
   flagship.
5. **Deployment.** Pages, Cloudflare Pages or Netlify. The CI currently builds and verifies but does
   not publish.
