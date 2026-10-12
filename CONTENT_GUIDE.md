# Content guide

The authoring contract. Page shape consistency is the product: it is what lets a reader skim forty
lessons the same way. Read this before writing a lesson.

## House style

1. **No dashes in prose or bullets.** Use commas, colons, semicolons or parentheses. Bullets use
   `*`. Ranges are written "20 to 30", never with a dash. The two exceptions are forced by tooling:
   file and folder names keep their `01-` and `I-` prefixes, because the site parses them for
   ordering, and markdown table delimiter rows need their `|---|`.
2. **Make a claim.** Every lesson argues something. If a paragraph could appear in any article on
   the topic, cut it.
3. **Numbers or nothing.** A statement about magnitude carries a number, a source and a date.
4. **Mechanism outlives tooling.** Explain how the thing works first. Name tools last, in their own
   subsection, with the caveat that the list will have changed.
5. **British spelling**, Oxford comma, sentence case headings.

## Page shape

````markdown
# Lesson Title

```text
Level: Intermediate
Prerequisites: The Product Loop
Estimated difficulty: 3/5 | Estimated time: 40 min
```

## The idea in one minute
## A picture
## How it really works
### What fills the box in 2026
## Run the numbers
## The AI layer
## Sector lens
## Remember this
## Try it
## Check yourself
````

Notes on the shape:

* The `text` block is parsed by the site for the difficulty badge, prerequisites and reading time.
  Keep the three field names exactly as written, and make `Level:` match the level the topic sits at
  in `config.yaml`. A lesson that says Expert inside a topic listed under Sectors prints both chips
  and reads as a mistake. `go run ./tools/levelcheck` enforces this, which matters because topics
  move between levels and the lessons do not follow on their own.
* **Topic folders carry no number.** Order comes from the `levels` lists in `config.yaml`, so a
  restructure is a config edit rather than a mass rename. Lesson files keep their `01-` prefixes,
  which the site does parse for ordering.
* The first paragraph after that block becomes the card blurb and the meta description, so it has
  to stand alone. Do not open with "In this lesson we will".
* `## Run the numbers` is included only where arithmetic decides the answer. A lesson about a
  ritual does not need a program.
* `## The AI layer` is required only where the lesson's subject is not already the model: topic I,
  the metrics topics, the sector topic. On an AI native lesson the whole page is the AI layer, and
  repeating the heading is padding. Where it does appear it must be specific: where a model fits,
  what it costs per unit of work, how it fails. "Nothing yet, and here is why" beats a hedge.
* `## Sector lens` is a table with one row per sector, linking to
  `product-management/XIII-product-by-sector`. Never restate market shape; state the delta.
* `## Try it` has three tasks and the third applies the lesson to the reader's own product.

## Where a file goes

The course page and the sidebar build themselves from the folder layout, so placement is the only
instruction the site needs.

| Put it here | It appears as |
|---|---|
| A folder listed under a level in `config.yaml` | A topic, under that level, in order |
| A folder **not** listed under any level | A card under **Build** |
| A loose `.md` file at the course root | A card under **Reference** |
| `README.md` in any folder | That folder's landing page |

So `projects/` is the Build shelf, and `ROADMAP.md`, `WORKFLOWS.md`, `REFERENCES.md` and
`GLOSSARY.md` are the Reference shelf, purely because of where they sit. Nothing registers them.

**Every link in `REFERENCES.md` must have been checked live before it is added**, and the page
carries the date it was checked. A source that cannot be reached from a plain request is left out
rather than quoted from memory, which is the same standard the lessons apply to numbers.

## Diagrams

Written as Mermaid, rendered in a hand sketched style. One diagram per section at most, and each
one answers exactly one question. If it needs a legend, it is two diagrams.

**Only these six classes exist.** They are injected by `assets/js/main.js`; a new name would need a
code change and would render unstyled.

| Class | Colour | Use it for |
|---|---|---|
| `neutral` | grey | people, market, demand, context |
| `compute` | blue | the work being done, the team, the system |
| `queue` | yellow | decisions, gates, stages, approvals |
| `io` | purple | external touchpoints: channels, buyers, partners, integrations |
| `memory` | green | stores and retained value: CRM, warehouse, money, the installed base |
| `warn` | red | risk, churn, failure, leakage |

The sidebar is a file tree: every level and every topic is a `<details>` with a disclosure
chevron, nested under indent guides, so a reader can collapse what they are not working on. Levels
start open. Nothing needs doing per topic; the structure comes from the folder layout.

Separately from the diagram palette, each **level** carries a colour, set in `assets/css/main.css`
as `--l-foundations` through `--l-sectors` and applied through a `data-level` attribute on the
level badge and the sidebar label. The ramp runs green, teal, blue, orange, crimson as difficulty
rises, and Sectors sits on violet because it is a different axis rather than a harder one. Adding a
level means adding one variable and one selector.

Six grammars cover almost everything in this library. Pick one before drawing.

1. **Loop** for anything that compounds.
2. **Funnel** for stage conversion.
3. **Two by two** for a strategy choice.
4. **Swimlane** for who owns what.
5. **Layer stack** for an architecture.
6. **System map** for data and money flow.

Syntax rules that keep diagrams from breaking:

* Quote every label that contains HTML: `A[":i-users: <b>Users</b><br/><small>1,000/wk</small>"]`.
* Node ids are plain alphanumerics. Never use `end` as an id.
* Assign classes in one statement per class: `class A,B compute`.
* Label an edge with `A -->|"38% drop"| B`.
* A logo token only renders if `static/logos/<name>.svg` exists. Check before using one, or it
  prints literally as `:name:` on the page. Run `go run ./tools/logos` to add more.
* Put one sentence under the diagram saying what to notice in it.

## Go blocks

Standard library only, `package main`, and they must build. `go run ./tools/gocheck -run <course>`
formats every block in place, vets it and executes the ones that need no arguments or network.
Print output that teaches something: a table, or a number next to the naive version of the number.

## The four checks

Run all four before committing. They exist because three hundred lessons cannot be proofread by eye.

```bash
go run ./tools/gocheck -run product-management   # format, vet and execute every Go block
go run ./tools/stylecheck product-management     # a dash left in prose or a bullet
go run ./tools/diagramcheck product-management   # logo tokens, class names, HTML in labels
go run ./tools/linkcheck -pending                # dangling links, grouped by unwritten folder
go run ./tools/configcheck                       # the two ways config.yaml truncates silently
go run ./tools/levelcheck                        # a lesson whose Level disagrees with config.yaml
```

`diagramcheck` catches the failures that render silently rather than loudly: a logo token with no
file prints as `:name:` on the page, a class name outside the six renders unstyled, and an
unbalanced `<b>` makes Mermaid drop the label. `linkcheck -pending` is the view to use while a
course is being written; forward references are tracked rather than forbidden, and the strict form
has to pass before a course is called complete.

Three gotchas that have each cost time already:

* **Never build into `public/` while `hugo server` is running.** This version renders the server to
  disk, so `hugo` and `rm -rf public` fight the process serving your pages: the site goes stale or
  404s and nothing in the log says why. Start the server with `--renderToMemory` and let
  `./check.sh` build to a temporary directory.

* **`config.yaml` fails silently in two ways, and `go run ./tools/configcheck` catches both.**
  A colon inside an unquoted value makes it a nested mapping, which breaks the build with a
  cryptic message; the house style uses colons where most writing uses dashes, so quote every
  `summary` and `description`. Worse, inside a flow mapping `{ a: 1, b: 2 }` a comma starts a new
  key, so `blurb: Start here, and learn the job.` silently keeps only "Start here" and the site
  builds happily with a truncated string. Use block mappings, never flow mappings.
* **Home page card summaries wrap to two lines at about 100 characters.** Longer than that and the
  cards stop lining up, which is what the length is tuned for.

One thing no tool can check: **the prose around a program must match what the program prints.**
Run the program, read its output, then write the paragraph. Three of the first five programs in
this library had narratives that the output contradicted, and every one of them was caught by
running it rather than by reading it.

## Definition of done

* [ ] `# Heading` present, and the `text` metadata block complete
* [ ] The idea in one minute stands alone and makes a claim
* [ ] At least one diagram, using only the six classes, with a caption line
* [ ] Every logo token resolves to a file in `static/logos/`
* [ ] Go blocks pass `gocheck -run`, and the prose matches what they actually print
* [ ] `diagramcheck` and `stylecheck` are clean
* [ ] Every metric links to `metrics-and-models` and is defined nowhere else
* [ ] The AI layer names a cost and a failure mode
* [ ] Sector lens covers the sectors that genuinely differ, linked
* [ ] Every number sourced and dated
* [ ] Remember this, Try it, Check yourself all present
* [ ] No dashes in prose or bullets
* [ ] Links resolve both on GitHub and on the site
* [ ] The topic `README.md` lesson table includes the new lesson
