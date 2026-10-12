# Prodloop

A working library on how modern AI products are actually built, published as a static site at
[product.kartikbhalerao.in](https://product.kartikbhalerao.in/).

| Course | What it covers |
|---|---|
| [`product-management/`](product-management/README.md) | The product function from first principles: strategy, discovery, delivery, the technical and design literacy it needs, the archetypes a product can be, the sectors it sells into, and how the job is run as an organisation. 16 topics. |
| [`metrics-and-models/`](metrics-and-models/README.md) | Every number an operator is judged on: what it measures, the formula, the denominator people get wrong, and a program that computes it. 7 topics. |
| `growth-engineering/` | Growth as a system rather than a list of tactics. *Coming soon.* |
| `gtm-engineering/` | Taking a product to a market: buyer, positioning, motion, pipeline, launch. *Coming soon.* |
| `revenue-operations/` | Operating cadence, the funnel as a data model, systems of record, process. *Coming soon.* |
| `ai-engineering/` | Building products with a model inside them, up to the product surface and no further. *Coming soon.* |

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
hugo server            # http://localhost:1313
```

Needs Hugo extended 0.126 or newer, because the site is generated from the course folders by a
content adapter rather than from `content/`. Go is needed only for the tools in `tools/`.

```bash
go run ./tools/gocheck -run product-management    # format, vet and run every Go block
go run ./tools/logos                              # revendor diagram icons (needs network)
```

## Repository layout

```
product-management/     a course: folders are topics, files are lessons
metrics-and-models/     a course
config.yaml             which courses are published, and the level each topic sits at
content/                the content adapter that turns the folders above into a site
layouts/ assets/        the site itself. Adding a course does not touch these.
data/seo.yaml           per page search metadata for high value lessons
static/logos/           the marks and glyphs diagrams draw by name
tools/                  gocheck, logos, seocheck
CONTENT_GUIDE.md        the authoring contract: page shape, diagram rules, definition of done
PLAN.md COURSE-MAP.md   the build plan and the full lesson map
```
