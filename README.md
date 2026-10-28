# Fathom

A working library on how modern AI products are actually built, published as a static site at
[product.kartikbhalerao.in](https://product.kartikbhalerao.in/).

A fathom is a unit for measuring depth. To fathom something is to understand it. This library tries
to do both: measure honestly, with numbers you can rerun, and explain the mechanism underneath.

| Course | What it covers |
|---|---|
| [`product-management/`](product-management/README.md) | Building products with a model inside them. The product function, what a model changes, discovery, behaviour specs, context and agents, evaluation, pricing, cost, trust and safety, the sectors, and every metric the job is judged on. 23 topics. |
| [`growth-engineering/`](growth-engineering/README.md) | Growth as a system rather than a list of tactics: the loops, the channels, the product surfaces that move a rate, and the instrumentation underneath. The product course owns every metric definition; this one links to them. Foundations and Basic complete, 7 topics of 19 written. |
| `gtm-engineering/` | Taking a product to a market: buyer, positioning, motion, pipeline, launch. *Coming soon.* |
| `revenue-operations/` | Operating cadence, the funnel as a data model, systems of record, process. *Coming soon.* |
| `ai-engineering/` | Reserved. Most of this material now lives inside the product course. *Coming soon.* |

Alongside the lessons sit four reference pages: a **roadmap** with four routes through the course,
**workflows** as runnable procedures, **references** to verified outside sources, and a
**glossary**. Start with the roadmap.

Courses are organised by **topic**: a course is a list of topics, a topic is a list of lessons.
Nothing is numbered as a module.

Every number quoted in a lesson carries its source and the date it was checked. Benchmarks rot
faster than mechanisms, and each lesson separates the two so you know which half to recheck.

## How a lesson is built

Every lesson follows the same shape, so forty of them can be read the same way.

1. **The idea in one minute.** The whole lesson in a few sentences, making a claim.
2. **A picture.** The mechanism, drawn. Press **Expand** on any diagram to open it full size.
3. **How it really works.** The precise version, with names and numbers.
4. **Run the numbers.** Where arithmetic decides the answer, a small Go program that uses only
   the standard library. You can rerun it with your own inputs.
5. **The AI layer.** What a model changes about this specific piece of work, including what it costs
   and how it fails.
6. **Sector lens.** How the same mechanism bends in enterprise software, fintech, ecommerce,
   marketplaces, devtools, healthtech and vertical software.
7. **Remember this**, **Try it** and **Check yourself.**

Diagrams are written as Mermaid and drawn in a hand sketched style. A diagram can name a tool by
its logo, and a tool named in a diagram is an example of what fills the box, not a recommendation.

## Read it

```bash
hugo server --renderToMemory    # http://localhost:1313
./check.sh                      # every gate, building to a temp directory
```

**Pass `--renderToMemory`.** This version of Hugo renders the dev server to `public/`, so any
ordinary `hugo` build while the server is running deletes and rewrites the files it is serving,
and the site goes stale or starts returning 404 with nothing in the log to explain it. Rendering
to memory makes `public/` irrelevant to the server, and `check.sh` builds to a temporary directory
for the same reason.

Needs **Hugo extended 0.146.0 or newer**, and the floor is hard rather than cautious: `layouts/`
uses the template system that shipped in 0.146.0, where partials live in `layouts/_partials` and
`baseof`, `home`, `page` and `section` sit at the root of `layouts`. An older Hugo finds none of
them, so the build fails on the first page with `partial "seo/head.html" not found` rather than
with anything about versions. The content adapter that generates the site from the course folders
instead of from `content/` needs 0.126, which is where the old and wrong floor came from.

Go is needed only for the tools in `tools/`. Two places pin the version and both have to agree with
this line: `.github/workflows/checks.yml`, and the `HUGO_VERSION` variable in the Cloudflare Pages
project, whose own default is older than the floor on every build image before v3.

```bash
go run ./tools/gocheck -run product-management     # format, vet and run every Go block
go run ./tools/stylecheck product-management      # a dash left in prose or a bullet
go run ./tools/diagramcheck product-management    # logo tokens, class names, HTML in labels
go run ./tools/linkcheck -pending                 # dangling links, by unwritten folder
go run ./tools/logos                              # revendor diagram icons (needs network)
```

## Repository layout

```
product-management/     the course
  <topic>/              a topic: README.md is its landing page, 01-*.md are lessons
  projects/             the Build shelf: things to construct, one folder each
  ROADMAP.md            the Reference shelf: any loose .md at the course root becomes a card
  WORKFLOWS.md          nine procedures, each with steps and what it produces
  REFERENCES.md         outside sources, every link checked live
  GLOSSARY.md           the terms this course uses
config.yaml             which courses are published, and the level each topic sits at
content/                the content adapter that turns the folders above into a site
layouts/ assets/        the site itself. Adding a course does not touch these.
data/seo.yaml           per page search metadata for high value lessons
static/logos/           the marks and glyphs diagrams draw by name
tools/                  gocheck, logos, seocheck
CONTENT_GUIDE.md        the authoring contract: page shape, diagram rules, definition of done
PLAN.md COURSE-MAP.md   the build plan and the full lesson map
```
