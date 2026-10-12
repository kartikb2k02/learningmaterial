# Course Map: AI Product

**Site:** product.kartikbhalerao.in · **Repo:** github.com/kartikb2k02/learningmaterial

Positioning after the revamp: **how modern AI products are actually built.** Not classic product
management with an AI chapter bolted on, and not an engineering course. The subject is the product
work that only exists because there is a model in the product: specifying behaviour you cannot
pin down, evaluating output you cannot assert on, pricing a feature whose cost scales with use,
and designing for an answer that is sometimes wrong.

Five courses. Lesson rows below are written as *the question it answers*, which is the same
convention the topic `README.md` tables use, so each table is copy and paste ready.

| Course | Folder | Topics | Lessons | Status |
|---|---|---|---|---|
| **AI Product Management** ⭐ | `ai-product-management` | 16 | ~107 | topics I and XIII partly written |
| Metrics and Economics | `metrics-and-economics` | 8 | ~51 | topic IV partly written |
| AI Growth | `ai-growth` | 8 | ~47 | announced |
| AI Go To Market | `ai-gtm` | 8 | ~46 | announced |
| AI Operations | `ai-operations` | 8 | ~46 | announced |

About **297 lessons** at full build. A publishable v1 is the flagship plus the metrics course,
about 158 lessons.

---

## What changes, and what survives

Everything written so far survives. The revamp replaces unwritten topics, so there is no rework.

| Already written | Fate |
|---|---|
| `product-management/I-the-product-function`, 6 lessons | **Unchanged.** Stays topic I. The job is the job |
| `product-management/XIII-product-by-sector`, README plus Enterprise | **Stays topic XIII.** Each sector lesson gains a full AI section instead of one table row |
| `metrics-and-models/IV-revenue-and-efficiency/01-cac-payback-and-ltv` | **Stays topic IV.** Lesson unchanged |
| `tools/` (gocheck, stylecheck, diagramcheck, linkcheck, logos), config, icon set, CI | Unchanged |

Four changes to make before writing resumes.

1. **Rename the course folder** `product-management` to `ai-product-management`, and
   `metrics-and-models` to `metrics-and-economics`. Two folder renames, two `config.yaml` entries,
   two `module.mounts` lines, and about six cross links.
2. **Retitle the site.** "Operator" no longer describes it. Proposed title **AI Product**, hero
   line *How AI products are actually built*.
3. **Drop the separate `ai-engineering` course.** It is absorbed into the flagship, because an AI
   course inside an AI library is a tautology.
4. **Rewire one prerequisite.** The sector README currently points at a classic
   `IX-product-archetypes`, which is replaced by `II-04 AI Product Shapes`.

### What is deliberately dropped

Honest about the cost of the reposition. Classic platform and API product management, classic
product archetypes such as marketplace liquidity and network density, and classic non generative
machine learning products (rankers, risk scores, forecasts) all lose their dedicated topics.

Marketplace and network mechanics survive where they matter, inside
`XIII-product-by-sector/04-marketplaces.md`. The rest is genuinely gone. If you want it back, it
belongs in a sixth course called `product-foundations` rather than diluting the flagship.

---

# 1. `ai-product-management` ⭐

> Building products with a model inside them. What a model changes about the product, how to
> specify behaviour that is not deterministic, how to know it works, how it should feel, what it
> costs to run, and how all of it bends by sector.

### I: The Product Function  `[Foundations]` ✅ written

*The job itself, before any model. Unchanged by the revamp.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | What a Product Manager Does | What are the three actual jobs, and what four things is this job not? |
| 02 | The Product Loop | Discover, decide, build, measure, learn: where do teams break the loop? |
| 03 | The Artifacts | What is each document for, and when is writing one pure cost? |
| 04 | Influence Without Authority | Where does the power come from when nobody reports to you? |
| 05 | How Product, Engineering and Design Split the Work | Where are the real boundaries, and what breaks at each seam? |
| 06 | What Good Looks Like | How is the job judged, and why is one quarter of outcomes close to noise? |

### II: What a Model Changes  `[Foundations]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | The Five Facts | Slow, costly, probabilistic, priced by the token, persuadable: what follows from each? |
| 02 | Deterministic and Probabilistic Products | What breaks in your spec, your QA, your support and your pricing? |
| 03 | The Capability Map | What can models do and not do as of today, and how do you keep the answer current? |
| 04 | AI Product Shapes | Copilot, agent, autonomous workflow, AI native app, infrastructure, embedded: which are you building? |
| 05 | Feature or Product | Where does value accrue in the stack, and what is the thin wrapper risk? |
| 06 | When Not to Use a Model | What is the deterministic alternative, and how do you argue for it? |
| 07 | The 2026 Landscape | Frontier labs, open weights, inference providers and tooling, on one map |

### III: Discovery for AI Products  `[Basic]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Demos Lie | Why do AI demos convert and AI products churn? |
| 02 | Find Tasks, Not Features | How do you decompose work into tasks a model could hold? |
| 03 | The Ceiling Test | Can the model do this at all, and how do you find out in a day? |
| 04 | Wizard of Oz and Concierge | How do you test demand with humans before any model? |
| 05 | Sizing an AI Opportunity | Volume times value times tolerance for error: how do you put numbers on it? |
| 06 | The Error Tolerance Question | What happens when it is wrong, and who absorbs that? |

### IV: Specs for Non Deterministic Systems  `[Basic]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Why a Pixel Spec Fails | How do you specify behaviour when the output is different every time? |
| 02 | The Behaviour Contract | Must, must never, should, may: how do you write one? |
| 03 | Writing the Rubric | How does quality get defined before it can be measured? |
| 04 | The Prompt Is the Spec | Who owns it, who reviews it, and how is it versioned? |
| 05 | Golden Examples and Counter Examples | How does the spec become a dataset? |
| 06 | The AI PRD | What does the document actually look like now? |

### V: Model Primitives  `[Basic]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Tokens and the Context Window | What is the budget you are spending, and what happens at the edges? |
| 02 | Latency That Users Feel | Time to first token, tokens per second, streaming, and why p95 is the number |
| 03 | Cost Per Request | What is the arithmetic every AI PRD needs and almost none contain? |
| 04 | Structured Output and Schemas | How do you get something a program can rely on? |
| 05 | Choosing a Model | Frontier or small, hosted or open weights: what is the real tradeoff? |
| 06 | Routing, Caching and Batching | What are the three cost levers, and what does each one cost you in quality? |
| 07 | The Non Determinism Problem | Temperature, seeds, reproducibility, and what it does to your test suite |

### VI: Context and Knowledge  `[Intermediate]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Context Engineering | What goes in the window, what it costs, and what gets lost in the middle? |
| 02 | Retrieval and Grounding | What does retrieval actually fix, and what does it not? |
| 03 | Chunking, Ranking and Recall | Which retrieval decisions are product decisions? |
| 04 | Citations and Provenance | How do you make an answer checkable by the person reading it? |
| 05 | Memory and Personalisation | Session, user and organisation memory as a product surface |
| 06 | Knowledge That Goes Stale | What is the half life of a wrong answer, and how do you invalidate? |
| 07 | When Finetuning Is the Answer | Rarely. What do you do instead, and what is the exception? |

### VII: Agents and Autonomy  `[Intermediate]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | The Autonomy Spectrum | Suggest, draft, act with approval, act: where should this product sit? |
| 02 | When a Workflow Beats an Agent | The most money saving lesson here. How do you tell? |
| 03 | Tools and Actions | What is the model allowed to touch, and who decided? |
| 04 | MCP and the Integration Layer | What does an agent native integration story look like? |
| 05 | Approval Gates and Human in the Loop | Where does a human belong, and what does their attention cost? |
| 06 | Blast Radius and Reversibility | How do you design for the wrong action rather than against it? |
| 07 | Multi Step Reliability | Why is 95% per step only 60% over ten steps, and what do you do? |
| 08 | Agent Identity and Delegated Permissions | Who is the agent acting as, and what may it do on their behalf? |

### VIII: AI Product UX  `[Intermediate]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | The Trust Budget | You get a fixed number of wrong answers. How do you spend them? |
| 02 | Streaming and Perceived Latency | How do you make slow feel fast without lying? |
| 03 | Showing Uncertainty | How do you communicate doubt without making the product feel broken? |
| 04 | Designing for Wrong Answers | Undo, edit, diff review, escalate: what has to exist? |
| 05 | Generative and Adaptive Interfaces | When should the interface itself be generated? |
| 06 | Voice and Multimodal Surfaces | What changes when there is no screen, or more than text? |
| 07 | Feedback Loops in the Interface | How does the UI feed the eval set? |

### IX: Evaluation and Quality  `[Advanced]` ★ flagship topic

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Why You Cannot Unit Test a Model | What replaces an assertion, and what does that cost? |
| 02 | Building the Dataset | Why thirty real cases beat five thousand synthetic ones |
| 03 | Graders: Code, Judge and Human | Which grader for which claim, and what is each one blind to? |
| 04 | A Judge Prompt That Agrees With Humans | How do you validate the judge before trusting its numbers? |
| 05 | The Statistics You Cannot Skip | Confidence intervals, repeated runs, and reading flipped cases |
| 06 | The Release Gate | What has to pass before a prompt, a model or a retrieval setting ships? |
| 07 | Online Evals and Reading Traces | What do you measure on live traffic, and how often do you read it? |
| 08 | Evals as the Product Artifact | Why the eval set, not the PRD, is the real specification |

### X: AI Pricing and Packaging  `[Advanced]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Why Seats Break | What happens when cost scales with usage and price does not? |
| 02 | The Value Metric for AI | Tasks, resolutions, outcomes or seats: what do you charge for? |
| 03 | Credits and Usage Pricing | How does a meter work, and how do you keep it from feeling hostile? |
| 04 | Outcome Based Pricing | The trend of the year. When is it right and when is it a trap? |
| 05 | Margin Under Load | What happens when the pricing model and the cost model disagree? |
| 06 | Free Tiers When Inference Costs Money | How much can you give away, and to whom? |
| 07 | Repricing as Model Costs Fall | Token prices drop every year. Who captures that? |

### XI: Trust, Safety and Guardrails  `[Advanced]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Prompt Injection as a Product Problem | Why is this not only a security issue, and what does product own? |
| 02 | The Guardrail Stack | Input, output and action guardrails: what does each actually catch? |
| 03 | Jailbreaks, Abuse and Rate Limits | How will this be abused in week one, and what is the ladder? |
| 04 | Redaction and Data Boundaries | What must never reach the model, and how do you prove it did not? |
| 05 | Content Policy for Generative Output | How do you write a policy, enforce it, and handle appeals? |
| 06 | Transparency Duties and the AI Act | What are you obliged to disclose, as of today, and to whom? |
| 07 | Incident Response for an AI Feature | What is the runbook when the model says something it should not have? |

### XII: Data Moats and Feedback Loops  `[Advanced]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Is Your Data Actually a Moat | What is the honest test, and how many companies pass it? |
| 02 | The Feedback Flywheel | How does usage become an eval set and then become quality? |
| 03 | Data Rights, Consent and Training | What may you use, and what must you promise not to? |
| 04 | The Thin Wrapper Critique | Where does value accrue when the model is somebody else's? |
| 05 | Switching Costs for AI Products | What actually holds a customer when the model is a commodity? |
| 06 | Build, Buy or Wait for the Model | How do you make a decision when next year's model may do it for free? |

### XIII: Product by Sector  `[Expert]` ✅ README and lesson 01 written

*Each lesson is a full product treatment of one market, now with its AI section promoted from a
table row to a section of its own. The README carries the cross sector comparison matrix that
every other course links to. Read `II-04 AI Product Shapes` first: shape usually matters more than
market.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Enterprise B2B SaaS ✅ | Why does the buyer want things the user will never see, and what does AI add to the questionnaire? |
| 02 | Fintech | What changes when being wrong moves money and a model may not decide alone? |
| 03 | Ecommerce and D2C | Why is this a margin business, and where does a model pay for itself per order? |
| 04 | Marketplaces | How do you serve two opposed customers, and what does matching by model change? |
| 05 | Devtools and Infrastructure | What happens when your users can read the source and your product writes code? |
| 06 | Healthtech and Regulated Care | How do you build around a clinician when a model may not diagnose? |
| 07 | Vertical SaaS | How do you win a small market completely, and why is AI the strongest case here? |

### XIV: Model Lifecycle and Drift  `[Expert]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | A Model Upgrade Is a Breaking Change | Why does a better model make your product worse, and how often? |
| 02 | Version Pinning and Migration | What do you pin, and what is the migration runbook? |
| 03 | Detecting Drift in Production | What signals move first when quality degrades? |
| 04 | Regression Suites for Prompts and Models | How do you keep a fix from being undone by the next change? |
| 05 | Deprecating a Model You Depend On | The provider retires it in ninety days. What now? |
| 06 | Vendor Risk and Multiple Providers | Is portability worth what it costs you in quality? |

### XV: AI Product Org and Process  `[Expert]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | The New Team Shape | PM, AI engineer, eval owner, designer: who owns what? |
| 02 | Prototyping as a Product Manager | What changes when you can build the thing yourself in an afternoon? |
| 03 | The AI Review Ritual | What is the checklist before an AI feature ships? |
| 04 | Trace Triage as an Operating Habit | Who reads production output every week, and what do they do with it? |
| 05 | Hiring for AI Product Work | What do you actually test for in an interview? |
| 06 | What to Centralise | Platform team or embedded: which parts of this should be shared? |

### XVI: Worked AI Products  `[Expert]`

*Each lesson runs one real build end to end: situation, numbers, options, decision, how it is
measured, and what would falsify it.*

| # | Lesson | What it works through |
|---|---|---|
| 01 | An In Product Copilot | Suggestion acceptance, trust budget, and the cost per accepted suggestion |
| 02 | An AI Support Assistant | Containment rate against satisfaction, and the escalation design |
| 03 | An AI Analyst Feature | Grounding on a warehouse, citations, and what happens when the number is wrong |
| 04 | An Agentic Workflow That Acts | Approval gates, blast radius, and the reliability maths over ten steps |
| 05 | A Model Upgrade Gone Wrong | A real regression, found late, and the postmortem |
| 06 | The AI Product Review | The checklist, runnable against any AI feature |

### `projects/`, Build

`01-eval-harness` · `02-cost-per-task-model` · `03-retrieval-quality-bench` ·
`04-trust-budget-calculator` · `05-agent-reliability-sim` · `06-ai-pricing-model` ·
`07-guardrail-test-set` · `08-drift-detector`

---

# 2. `metrics-and-economics`

> Every number an AI product is judged on, defined once. Classic revenue arithmetic first, because
> it still decides whether the business works, then the quality, cost and agent metrics that only
> exist because there is a model in the product.
>
> **This course is the definition site. No other course defines a metric; they link here.**

### I: How to Read a Metric  `[Foundations]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | What a Metric Is For | Which decision does this number change, and if none, why is it on a dashboard? |
| 02 | The Denominator Problem | Why is conversion rate three different numbers, and which did you mean? |
| 03 | Rates, Ratios and Averages | Why is the average nearly always the wrong summary? |
| 04 | Simpson's Paradox and Mix Shift | How can every segment improve while the total gets worse? |
| 05 | Leading, Lagging and Vanity | How do you tell a predictive number from a flattering one? |
| 06 | The North Star and Its Critics | What is a single top metric good for, and where does it mislead? |

### II: Funnel and Acquisition  `[Basic]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Funnel Mechanics | How does loss compound across stages, and where is the leverage? |
| 02 | Traffic, Visits and Sessions | Why does every tool report a different number for the same thing? |
| 03 | Signup and Conversion Rate | What counts, over what denominator, in what window? |
| 04 | Acquisition Cost: Blended and Paid | What belongs in the numerator, and why is blended misleading? |
| 05 | Attribution | First touch, last touch, multi touch: in what way is each wrong? |
| 06 | Cost Per What | Translating between the cost per thousand, click, lead and action |

### III: Retention and Cohorts  `[Basic]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | What Retention Means | Which of the three common definitions are you using? |
| 02 | Building a Cohort Table | How do you build the one chart that tells the truth about a product? |
| 03 | The Curve and the Flattening | Does it flatten, and why does that decide whether you have a business? |
| 04 | Churn: Logo and Revenue | Why do these move in opposite directions? |
| 05 | Engagement Ratios | What does the daily over monthly ratio hide? |
| 06 | Resurrection and Dormancy | How big is the segment everyone forgets, and what is it worth? |

### IV: Revenue and Efficiency  `[Intermediate]` ✅ lesson 01 written

| # | Lesson | The question it answers |
|---|---|---|
| 01 | CAC Payback, and LTV ✅ | Which assumptions are load bearing, and why is payback the honest number? |
| 02 | MRR, ARR and the Movements | How do new, expansion, contraction and churn build the waterfall? |
| 03 | Net and Gross Revenue Retention | What do they measure, and how are they gamed? |
| 04 | Gross Margin for Software | What belongs in cost of revenue, inference included? |
| 05 | The Efficiency Metrics | Magic number, burn multiple, Rule of 40: what does each one test? |
| 06 | The Ratio Is Not a Strategy | What can lifetime value over acquisition cost actually tell you? |

### V: AI Quality Metrics  `[Intermediate]`

*The metrics that did not exist five years ago. Most AI products measure none of these.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Eval Pass Rate and Coverage | What does a pass rate mean, and what fraction of real behaviour does your set cover? |
| 02 | Groundedness and Citation Accuracy | How do you measure whether an answer is supported by its sources? |
| 03 | Task Success on End State | Why grade the outcome rather than the path, and how? |
| 04 | Acceptance Rate and Edit Distance | The copilot metrics: what fraction is kept, and how much is rewritten? |
| 05 | Intervention and Escalation Rate | How often does a human have to step in, and what does that cost? |
| 06 | Containment and Deflection | The support metric everyone quotes. What does it hide? |
| 07 | Retry, Regeneration and Abandonment | What do users do when the answer is bad, and how do you count it? |
| 08 | Measuring Trust | Can you measure it at all, and what are the honest proxies? |

### VI: AI Cost and Margin  `[Advanced]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Cost Per Task | The new line item. How do you compute it for a real feature? |
| 02 | Tokens Per Task and Context Efficiency | Where does the context budget go, and what is waste? |
| 03 | Inference in Cost of Revenue | What does it do to gross margin, and who books it today? |
| 04 | The Routing and Caching Savings Model | How much do the three cost levers actually save, and at what quality cost? |
| 05 | Margin Under Engagement Growth | What happens when your best cohort is also your most expensive? |
| 06 | The Cost of Quality | What do evals, judge calls and human review cost per task? |
| 07 | A Worked AI Margin Model | One feature, from token count to gross margin, end to end |

### VII: Agent and Task Metrics  `[Advanced]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Steps Per Task and Loop Depth | What does a rising step count tell you before quality drops? |
| 02 | Tool Error and Recovery Rate | How often do actions fail, and does the agent recover? |
| 03 | Autonomy Ratio | What share of actions happen without a human, and should it rise? |
| 04 | Time to Task Completion | Why is this the metric users actually feel? |
| 05 | Cost and Latency Percentiles | Why does the mean hide the thing that makes people leave? |
| 06 | Reliability as a Compound Number | How do you turn per step rates into a number you can promise? |

### VIII: Benchmarks, Drift and Their Limits  `[Expert]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Where Benchmarks Come From | Why is most of what gets quoted unsourced, and how do you check one? |
| 02 | Public Model Benchmarks | What does a leaderboard score predict about your product? Usually little |
| 03 | Your Eval Set Is the Only Benchmark | How do you build the one number that means something to you? |
| 04 | Measuring Drift Across Versions | How do you compare two models on your own work, fairly? |
| 05 | Benchmarks by Sector | The table, every number sourced and dated |
| 06 | Reading a Public Filing | How do you get real figures instead of conference slides? |

### `projects/`, Build

`01-cohort-retention-table` · `02-cac-payback-calculator` · `03-arr-waterfall` ·
`04-cost-per-task-model` · `05-eval-pass-rate-with-intervals` · `06-agent-reliability-compounder`

---

# 3. `ai-growth`

> Distribution for AI products, where the usual playbook breaks twice: search traffic is being
> replaced by answers, and every free user now costs real money.

| Topic | Level | What it covers |
|---|---|---|
| `I-how-ai-products-grow` | Foundations | The loops that work for AI products, and the ones that stopped |
| `II-distribution-after-search` | Basic | Answer engines, citation as the new ranking, and what is left of content |
| `III-activation-and-time-to-first-output` | Intermediate | The empty prompt box problem, templates, and first useful result |
| `IV-retention-and-habit` | Intermediate | Why AI retention curves fall steeply, and what flattens them |
| `V-free-tiers-when-inference-costs-money` | Advanced | Giving away a marginal cost, abuse, and the economics of a free user |
| `VI-artifact-and-sharing-loops` | Advanced | When the output is the marketing, and how to design for it |
| `VII-the-growth-operating-system` | Expert | The experiment machine, when every experiment also changes cost |
| `VIII-growth-by-sector` | Expert | One delta lesson per sector |

### 4. `ai-gtm`

> Selling a product that is sometimes wrong, to buyers who have been promised AI by everyone.

| Topic | Level | What it covers |
|---|---|---|
| `I-market-and-buyer-for-ai` | Foundations | Who buys AI, who blocks it, and the new committee member |
| `II-positioning-against-we-have-ai-too` | Basic | Differentiating when every competitor shipped the same feature |
| `III-pilots-pocs-and-production` | Intermediate | The pilot trap, and what moves a POC into a contract |
| `IV-the-ai-security-review` | Intermediate | Subprocessors, training opt out, residency, and the model questionnaire |
| `V-motions-and-channels` | Advanced | What acquisition cost per deal looks like when the product sells itself unevenly |
| `VI-launch-and-expectation-management` | Advanced | Launching something probabilistic without overpromising |
| `VII-worked-gtm-plans` | Expert | Zero to one enterprise AI, and a self service AI launch |
| `VIII-gtm-by-sector` | Expert | One delta lesson per sector |

### 5. `ai-operations`

> Running an AI product: the traces, the costs, the vendors, the policy and the obligations.

| Topic | Level | What it covers |
|---|---|---|
| `I-why-operations-exists` | Foundations | What breaks without it, in an AI product specifically |
| `II-operating-cadence-for-ai-teams` | Basic | The weekly trace review, the cost review, the eval review |
| `III-ai-observability` | Intermediate | Traces, spans, token accounting and what to alert on |
| `IV-data-consent-and-retention` | Intermediate | What you keep, for how long, and what you promised |
| `V-vendor-and-model-management` | Advanced | Subprocessors, rate limits, fallbacks and contractual exposure |
| `VI-governance-policy-and-the-ai-act` | Advanced | Risk classification, documentation duties, and who signs |
| `VII-finance-ops-for-usage-pricing` | Expert | Metering, invoicing, overage, revenue recognition and bad debt |
| `VIII-operations-by-sector` | Expert | One delta lesson per sector |

---

## Cross course rules

1. **Metrics** are defined only in `metrics-and-economics`. Every other lesson links.
2. **Sector market shape** is defined only in `ai-product-management/XIII-product-by-sector`.
   Other courses' sector topics carry function deltas only.
3. **Product shape is read before market.** `II-04 AI Product Shapes` is a prerequisite for every
   sector lesson.
4. **Evals are owned by topic IX.** Other lessons reference the method rather than restating it.
5. **Cost arithmetic** lives in `metrics-and-economics/VI-ai-cost-and-margin`. Product lessons
   quote results and link.
6. Every course `README.md` ends with a **How it connects to the other courses** section.

## Build order

| Phase | Deliverable | Lessons |
|---|---|---|
| 0 | The renames, the retitle, and the prerequisite rewiring | 0 |
| 1 | Flagship topics **II, III, IV, V**: what a model changes, discovery, specs, primitives | ~26 |
| 2 | Flagship topics **VI, VII, VIII**: context, agents, UX | ~22 |
| 3 | Flagship topic **IX** evals, plus `metrics-and-economics` **V** quality metrics | ~16 |
| 4 | Flagship topics **X, XI, XII**: pricing, guardrails, moats | ~20 |
| 5 | `metrics-and-economics` **I, II, III, IV, VI, VII, VIII** | ~43 |
| 6 | Flagship **XIII** remaining six sectors, plus **XIV, XV, XVI** | ~24 |
| 7 | `ai-growth`, then `ai-gtm`, then `ai-operations` | ~139 |

Phases 1 to 4 plus 6 give a complete flagship, which with the written topic I is a real site on its
own. Phase 3 is the one to reach quickly: evaluation is the topic the whole positioning rests on.
