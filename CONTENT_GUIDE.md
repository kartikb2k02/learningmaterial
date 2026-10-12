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
  Keep the three field names exactly as written.
* The first paragraph after that block becomes the card blurb and the meta description, so it has
  to stand alone. Do not open with "In this lesson we will".
* `## Run the numbers` is included only where arithmetic decides the answer. A lesson about a
  ritual does not need a program.
* `## The AI layer` must be specific: where a model fits, what it costs per unit of work, how it
  fails. "Nothing yet, and here is why" is a better answer than a hedge.
* `## Sector lens` is a table with one row per sector, linking to
  `product-management/XIII-product-by-sector`. Never restate market shape; state the delta.
* `## Try it` has three tasks and the third applies the lesson to the reader's own product.

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
```

`diagramcheck` catches the failures that render silently rather than loudly: a logo token with no
file prints as `:name:` on the page, a class name outside the six renders unstyled, and an
unbalanced `<b>` makes Mermaid drop the label. `linkcheck -pending` is the view to use while a
course is being written; forward references are tracked rather than forbidden, and the strict form
has to pass before a course is called complete.

Two gotchas that have each cost time once already:

* **A colon inside an unquoted YAML value breaks `config.yaml`.** Because the house style uses a
  colon where most writing would use a dash, every `summary` and `description` has to be quoted.
  Run `hugo config` after any config edit; the failure is silent until the next build.
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
