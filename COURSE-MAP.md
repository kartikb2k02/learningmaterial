# Course Map — every topic, every lesson

Target repo: **`github.com/kartikb2k02/learningmaterial`** (private, empty — Hugo shell goes in at Phase 0).
Structure rules, lesson template and build phases: see [`PLAN.md`](PLAN.md).

**This file is the source for the topic `README.md` tables.** Each lesson row below is written as
*the question it answers* — same convention the reference repo uses. When you create a topic
folder, copy its table straight into that folder's `README.md`.

Six courses · ~6 lessons per topic · **~327 lessons at full build.** `product-management` is
the flagship and a complete v1 on its own.

| Course | Folder | Topics | Lessons | Phase |
|---|---|---|---|---|
| **Product Management** ⭐ | `product-management` | 16 | ~95 | 1 |
| Metrics & Models | `metrics-and-models` | 7 | ~41 | 1 |
| Growth Engineering | `growth-engineering` | 8 | ~47 | 2 |
| GTM Engineering | `gtm-engineering` | 8 | ~47 | 3 |
| Revenue Operations | `revenue-operations` | 8 | ~47 | 4 |
| AI Engineering | `ai-engineering` | 9 | ~50 | 5 |

**Sectors are not a separate course** — they live inside the course whose mechanism they modify.
`product-management/XIII-product-by-sector` is the primary home; see §7.

---

# 1. `product-management` — Product  ⭐ flagship course

> The product function from first principles, the technical and design literacy it needs, the
> archetypes a product can be, the sectors it gets sold into, and how it's run as an organisation.

**16 topics · ~95 lessons.** This is now the largest course by design — comparable to the
reference repo's `inference-engineering` (14 topics). Topics marked **★** are new; **◆** moved here.

### I — The Product Function  `[Foundations]`
*What the job actually is, before any technique.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | What a Product Manager Does | What are the three actual jobs, and what four things is this job not? |
| 02 | The Product Loop | Discover → decide → build → measure → learn: where do teams break the loop? |
| 03 | The Artifacts | PRD, one-pager, roadmap, brief, spec — what is each for, and when is it waste? |
| 04 | Influence Without Authority | Who do you need, what does each of them want, and how do decisions really get made? |
| 05 | How Product, Engineering and Design Split the Work | Where are the real boundaries, and what breaks at each seam? |
| 06 | What Good Looks Like | How is the job evaluated, and what is the difference between output and outcome? |

### II — Product Strategy  `[Foundations]` ★
*The biggest gap in most product curricula: strategy taught as a distinct thing from discovery.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Strategy, Vision and Roadmap | What is each one, and what do you get when you conflate them? |
| 02 | Where to Play and How to Win | Which two choices actually constitute a strategy? |
| 03 | Moats for Software | Switching cost, data, network, scale, brand — which are real and which are stories? |
| 04 | The Strategy Document | How do you write one page that states the bets *and* what you are not doing? |
| 05 | Build, Partner or Buy | How do you make the strategic version of this decision, with numbers? |
| 06 | When Strategy Should Change | What are the signals — and what does changing too often cost? |

### III — Discovery and Problem Framing  `[Basic]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Finding a Problem Worth Solving | Which demand signals mean something, and which are noise? |
| 02 | Customer Interviews That Aren't Useless | Which questions produce facts, which produce flattery, and how many people is enough? |
| 03 | Jobs, Alternatives and Doing Nothing | What do they do today — including nothing — and why would they stop? |
| 04 | Sizing a Problem | How many people, how often, and how much does it cost them now? |
| 05 | Writing the Problem Statement | How do you write one page that is falsifiable and contains no solution? |
| 06 | Killing Ideas Early | What is the riskiest assumption, and how do you test it before building? |

### IV — Definition and Delivery  `[Basic]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Writing a PRD That Gets Built | Which sections change what gets built, and which are read by nobody? |
| 02 | Scoping and the First Version | What does "minimum" actually mean, and what does it never mean? |
| 03 | Prioritisation That Survives Contact | RICE, WSJF and the rest — which is a tool and which is theatre? |
| 04 | The Roadmap Problem | How do you commit enough to be trusted without lying about dates? |
| 05 | Working With Engineering | Estimates, tradeoffs, tech debt — and when should you let them decide? |
| 06 | Launch Readiness | What has to be true before it ships: docs, support, pricing, legal, rollback? |

### V — Technical Foundations for PMs  `[Basic]` ★
*Enough systems literacy to not be dangerous. Deliberately stops short of engineering.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | How a Request Actually Travels | Client, API, service, database, cache: what is the mental model? |
| 02 | APIs for PMs | REST, webhooks, events, idempotency, rate limits — how does API shape constrain the product? |
| 03 | Data Models: Why Schema Is Product | What does a schema permit and permanently forbid? |
| 04 | Latency, Caching and the Cost of Slow | Where does time go, and which fixes are cheap? |
| 05 | Auth, Identity and Permissions | SSO, roles, tenancy: why are the features enterprise asks for so expensive? |
| 06 | What "It's Hard" Means | What are the five reasons engineers say no, and which are negotiable? |

### VI — Design and User Experience  `[Intermediate]` ★

| # | Lesson | The question it answers |
|---|---|---|
| 01 | What Designers Actually Do | What is the craft, and what are the three ways PMs get in its way? |
| 02 | Information Architecture and Navigation | Why is naming and grouping the hardest UI problem? |
| 03 | Interaction Patterns and When to Break Them | When does a convention cost more than it saves? |
| 04 | Running a Design Critique | How do you give feedback that improves the work rather than redirects it? |
| 05 | Usability Testing on a Budget | What can five users in one afternoon actually tell you? |
| 06 | Accessibility | What is required, what is decent, and why is this not optional? |

### VII — Analytics and Experimentation  `[Intermediate]`
*Every formula links to `metrics-and-models`.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Instrumentation Before Features | What event taxonomy has to exist before the feature ships? |
| 02 | Choosing What to Measure | What is the one primary metric, and what are the guardrails and counter-metric? |
| 03 | Reading a Funnel | Where is the drop-off actually, versus where does it look like it is? |
| 04 | A/B Testing and When It Lies | Power, duration, peeking, novelty — when is the result not real? |
| 05 | When You Cannot A/B Test | Holdouts, switchbacks, before/after, qualitative: what to do with low traffic? |
| 06 | Deciding With Ambiguous Data | How do you write a decision memo when the data does not settle it? |

### VIII — Pricing and Packaging  `[Intermediate]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | What Pricing Actually Is | Why is pricing value capture rather than cost-plus? |
| 02 | The Value Metric | What do you charge *for* — and why is this the single biggest pricing choice? |
| 03 | Packaging and Tiers | How do good/better/best, feature fences and the enterprise tier actually work? |
| 04 | Price Discovery | Van Westendorp, conjoint, WTP interviews — what does each tell you and where does it fail? |
| 05 | Changing Prices | Grandfathering, migration, communication: how do you raise price without a revolt? |
| 06 | Discounting and Floors | What does a discount do to your win rate, your margin and your next renewal? |

### IX — Product Archetypes  `[Advanced]` ★
*The **shape** of a product usually determines more than the sector it sells into. Read this
before the sector topic — a fintech marketplace behaves like a marketplace first.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Archetype Beats Sector | Why do two fintechs with different shapes have nothing in common? |
| 02 | The SaaS Tool | Seats, workflows, admin surface: the default archetype and its assumptions |
| 03 | The Marketplace | Liquidity, cold start, which side to subsidise, disintermediation, take rate |
| 04 | The Network Product | Density and clustering — and why most claimed network effects are not one |
| 05 | The Consumer App | Attention, habit, and the economics of free |
| 06 | Vertical SaaS | Depth in one industry, workflow replacement, and the services gravity problem |
| 07 | Embedded, B2B2C and Open Core | What changes when your customer's customer is the user, or the code is the funnel? |

### X — Platform and API Products  `[Advanced]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | When to Become a Platform | What does platform-ness cost, and what has to be true first? |
| 02 | Designing an API as a Product | What makes developer experience a product decision rather than a doc problem? |
| 03 | Pricing an API | Per call, per seat, per outcome — which matches which usage shape? |
| 04 | Ecosystems and Integrations | Build, partner or marketplace: how do you choose, and who maintains it? |
| 05 | Internal Platforms | What changes when the customer is another team and cannot churn? |
| 06 | Platform Failure Modes | No first-party use, breaking changes, abandoned SDKs — how do these start? |

### XI — Data and ML Products  `[Advanced]` ★
*Non-generative ML — still most of the ML in production, and the part `ai-engineering` doesn't cover.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | When a Model Is the Product | What changes when the output is a prediction rather than a feature? |
| 02 | Ranking and Recommendation | Which product decisions are hidden inside a ranker? |
| 03 | Scoring and Risk Models | Fraud, credit, churn — why is the threshold a product decision, not a data-science one? |
| 04 | Forecasting and Planning Products | Why do users reject accurate forecasts, and what do you do about it? |
| 05 | Cold Start and the Feedback Loop | How does a model that shapes its own training data go wrong? |
| 06 | When a Heuristic Wins | How do you tell, before spending six months? |

### XII — Trust, Safety and Abuse  `[Advanced]` ★
*Every product that holds money, content or identity gets abused. Rarely taught, always needed.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Every Product Gets Abused | What are the five vectors, and which apply to your archetype? |
| 02 | Fraud as a Product Problem | Why is this not a fraud-team problem, and what does product own? |
| 03 | Content Moderation and Policy | How do you write a policy, enforce it, and run appeals? |
| 04 | Identity and Verification | KYC, age assurance, proof-of-human: what does each cost in conversion? |
| 05 | The Abuse Metrics | Prevalence, enforcement precision and recall, appeal rate |
| 06 | The Enforcement Ladder | Warn, limit, suspend, ban — how do you design the escalation? |

### XIII — Product by Sector  `[Expert]` ◆
*The sector block, moved here from a standalone course. Each lesson is a **full product treatment**
of one market — not a thin delta page. The `README.md` holds the cross-sector comparison matrix
that every other course links to.*

| # | Lesson | What it covers |
|---|---|---|
| — | `README.md` | **The matrix:** buyer, cycle length, regulation, data constraints, margin shape and abuse surface for all seven sectors on one screen |
| 01 | Enterprise B2B SaaS | Procurement, security review, the admin surface, SSO/SCIM/audit logs, and the buyer-is-not-the-user problem |
| 02 | Fintech | Regulation as a roadmap constraint, money movement, reconciliation, KYC/AML, and why failure here is financial rather than annoying |
| 03 | Ecommerce and D2C | Catalogue, inventory truth, checkout, returns, margin per order, seasonality |
| 04 | Marketplaces | Two-sided liquidity, cold start, trust and disputes, disintermediation, take rate as a product lever |
| 05 | Devtools and Infrastructure | The developer as buyer, docs as product, DX, self-serve to enterprise, open core |
| 06 | Healthtech and Regulated Care | Clinical risk, HIPAA, designing around a clinician's workflow, buyer vs payer vs patient |
| 07 | Vertical SaaS | Depth in one industry, replacing a workflow rather than a tool, services gravity, small TAM with high ACV |

### XIV — Lifecycle, Migrations and Sunsets  `[Expert]` ★

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Zero to One vs One to N | Why is it a different job, with different skills and different metrics? |
| 02 | The Second Product Problem | Why do most second products fail, and what makes the exceptions work? |
| 03 | Migrations as Product Work | How do you move users off the old thing without losing them? |
| 04 | Versioning and Breaking Changes | When may you break something, and what do you owe people when you do? |
| 05 | Sunsetting a Feature | What is the sequence, the comms, and the long tail you will carry anyway? |
| 06 | Legacy and the Cost of Carrying It | How do you price the decision to keep supporting something? |

### XV — Product Operations and the Org  `[Expert]` ★

| # | Lesson | The question it answers |
|---|---|---|
| 01 | What Product Ops Is | What does it own, and at what size do you need it? |
| 02 | Feedback Aggregation | How do thousands of signals become a short, defensible list? |
| 03 | Research Ops and the Insight Library | How do you stop re-learning the same thing every year? |
| 04 | The Product Review and the Council | Which rituals work, and what makes them degrade? |
| 05 | PM Levels, Archetypes and Hiring | APM to Head of Product: what actually changes at each level? |
| 06 | Managing PMs and Your Own Career | How is the work judged, and what compounds? |

### XVI — Worked Product Decisions  `[Expert]`
*Each lesson runs one real decision: situation → numbers → options → decision → measurement → what would falsify it.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Should We Build This? | A full build / buy / partner / don't decision, worked with arithmetic |
| 02 | The Feature Nobody Used | A post-mortem with the numbers, and the signal that was there beforehand |
| 03 | A Pricing Change, End to End | From the model to the migration email to the 90-day readout |
| 04 | A Platform Bet, End to End | Committing to an API, and what it cost for two years |
| 05 | Entering a New Sector | One market entry, worked — using the archetype and sector topics together |
| 06 | The Product Review | The ritual, written as a checklist you can run on any proposal |

### `projects/` — Build
`01-prd-and-decision-memo` · `02-strategy-one-pager` · `03-funnel-instrumentation-plan` ·
`04-prioritisation-model` (Go) · `05-pricing-model` (Go) · `06-marketplace-liquidity-sim` (Go) ·
`07-threshold-and-abuse-model` (Go) · `08-sector-entry-brief`

---

# 2. `metrics-and-models` — Metrics

> Every number an operator is judged on: what it measures, the formula, the denominator people
> get wrong, what good looks like, and a program that computes it.
>
> **This course is the single source of truth. No other course defines a metric — they link here.**

### I — How to Read a Metric  `[Foundations]`
*Read this before any specific metric. It prevents most of the errors in the other six topics.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | What a Metric Is For | Which decision does this number change — and if none, why is it on the dashboard? |
| 02 | The Denominator Problem | Why is "conversion rate" three different numbers, and which one did you mean? |
| 03 | Rates, Ratios and Averages | Why is the average almost always the wrong summary, and what replaces it? |
| 04 | Simpson's Paradox and Mix Shift | How can every segment improve while the total gets worse? |
| 05 | Leading, Lagging and Vanity | How do you tell a predictive metric from a flattering one? |
| 06 | The North Star and Its Critics | What is a single top metric good for, and where does it actively mislead? |

### II — Funnel and Acquisition  `[Basic]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Funnel Mechanics | How does loss compound across stages, and where is the real leverage? |
| 02 | Traffic, Visits and Sessions | Why does every tool report a different number for the same thing? |
| 03 | Signup and Conversion Rate | What counts as a conversion, and over what denominator and window? |
| 04 | CAC: Blended vs Paid | What is the formula, what belongs in the numerator, and why is blended CAC misleading? |
| 05 | Channel Attribution | First touch, last touch, multi-touch — in what way is each of them wrong? |
| 06 | Cost Per What | CPM, CPC, CPL, CPA: how do you translate between them and sanity-check a quote? |

### III — Retention and Cohorts  `[Intermediate]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | What Retention Means | Which of the three common definitions are you using, and does your tool agree? |
| 02 | Building a Cohort Table | How do you build the one chart that tells the truth about a product? |
| 03 | The Retention Curve and the Flattening | Does the curve flatten — and why does that one question decide whether you have a business? |
| 04 | Churn: Logo vs Revenue | Why do these two move in opposite directions, and which should you report? |
| 05 | DAU, MAU and Engagement Ratios | What does the DAU/MAU ratio hide, and when is it meaningless? |
| 06 | Resurrection and Dormancy | How big is the segment everyone forgets, and what is it worth? |

### IV — Revenue and Efficiency  `[Intermediate]`
*The golden lesson (Phase 0) is `01` here — it exercises every element of the template.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | CAC Payback, and LTV | What are the formulas, which assumptions are load-bearing, and what is the honest version? |
| 02 | MRR, ARR and the Movements | How do new, expansion, contraction and churn build the revenue waterfall? |
| 03 | NDR, GRR and What Good Looks Like | What do they measure, how are they gamed, and what are real benchmarks? |
| 04 | Gross Margin for Software | What belongs in cost of revenue — including inference cost? |
| 05 | The Efficiency Metrics | Magic number, burn multiple, Rule of 40: what does each actually test? |
| 06 | LTV/CAC Is Not a Strategy | What can the ratio tell you, and what do people wrongly conclude from it? |

### V — Sales and Pipeline Math  `[Advanced]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Pipeline Coverage | Why is this the only forecasting number that works early, and what multiple is right? |
| 02 | Stage Conversion and Win Rate | What is a "stage", and how do you compute win rate without flattering yourself? |
| 03 | Sales Cycle Length | Why does the average cycle lie, and what should you report instead? |
| 04 | ACV, ASP and Deal Size Distribution | Why does the distribution matter more than the mean? |
| 05 | Quota, Capacity and Ramp | How many reps does a number require, and when do they have to start? |
| 06 | Forecasting Without Lying | Commit, best case, pipeline — and the three ways every forecast breaks |

### VI — Unit Economics With AI Costs  `[Advanced]`
*The newest and least-written-about part of operator maths.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Cost Per Task | What is the new line item, and how do you compute it for a real feature? |
| 02 | Token Costs in a P&L | Where do inference costs land, and what do they do to gross margin? |
| 03 | Margin Under Growth | What happens when usage scales faster than price? |
| 04 | Pricing Models vs Cost Models | Seats, credits, usage — what is the mismatch risk and how is it hedged? |
| 05 | The Cost of Quality | What do evals, human review and escalation actually cost per task? |
| 06 | A Worked Margin Model | One AI feature, from token count to gross margin, end to end |

### VII — Benchmarks and Their Limits  `[Expert]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Where Benchmarks Come From | Why is most of what gets quoted unsourced, and how do you check one? |
| 02 | Benchmarks by Sector | The table — every number sourced and dated |
| 03 | Stage-Adjusted Benchmarks | Why is a seed-stage number useless to a Series C company? |
| 04 | Reading a Public Filing | How do you get real benchmarks from a 10-K instead of a blog post? |
| 05 | Building Your Own Baseline | Why is your own trailing number the only benchmark that matters? |

### `projects/` — Build (all Go, stdlib only)
`01-cohort-retention-table` · `02-cac-payback-calculator` · `03-arr-waterfall` · `04-pipeline-coverage-model` · `05-ai-margin-model`

---

# 3. `growth-engineering` — Growth

> Growth as a system rather than a list of tactics: the loop, the model, the experiment machine,
> and the four places compounding actually comes from.

### I — How Growth Works  `[Foundations]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Growth Is a System, Not Tactics | Why do tactic lists stop working, and what replaces them? |
| 02 | The Four Loops | Content, viral, paid and sales loops — what makes each one compound or not? |
| 03 | The Growth Model | What is the equation your business actually runs on, and how do you write yours? |
| 04 | Growth vs Marketing vs Product | Who owns which lever, and what happens when nobody does? |
| 05 | Where Growth Comes From | Acquisition, activation, retention, monetisation: where is the leverage really? |
| 06 | The Ceiling Problem | Why does every channel saturate, and how early can you see it coming? |

### II — Acquisition  `[Basic]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Channel–Product Fit | Why will most channels never work for your product, and how do you tell in advance? |
| 02 | SEO and Content as a Loop | What makes content compound instead of decay — and what changed with AI answers? |
| 03 | Paid Acquisition Mechanics | Auctions, bids and LTV-bound spending: how much can you actually afford to pay? |
| 04 | Virality and Referral | What is k-factor, and why is yours almost certainly below 1? |
| 05 | Outbound and Communities | When does outbound work, and when is a community a channel rather than a nice-to-have? |
| 06 | Choosing Two Channels | What is the arithmetic case for focus over breadth? |

### III — Activation and Onboarding  `[Intermediate]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | What Activation Is | How do you define the moment instead of guessing at it? |
| 02 | Finding Your Activation Event | How do you use correlation-to-retention to find the real one? |
| 03 | The Setup Tax | What does each additional step cost you, in users, as arithmetic? |
| 04 | Time to Value | How do you shorten it without faking the value? |
| 05 | Onboarding Patterns | Checklists, empty states, templates, guided setup, concierge — which fits which product? |
| 06 | Activation for Teams | What changes when the buyer is not the user, and the value needs two people? |

### IV — Retention and Engagement  `[Intermediate]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Retention Is the Whole Business | Why does every other metric depend on this curve? |
| 02 | Habit, Frequency and the Natural Rate | Why should some products not be daily, and what is yours? |
| 03 | Why Users Leave | What are the four causes, and which of them can you actually fix? |
| 04 | Resurrection and Winback | Who comes back, what brings them back, and what is it worth? |
| 05 | Engagement: Depth vs Breadth | Which one predicts retention for your product shape? |
| 06 | The Retention Roadmap | How do you turn a flat curve into a sequenced plan? |

### V — Monetisation and Expansion  `[Advanced]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | The Free-to-Paid Boundary | What do you give away, and what does giving away too much cost? |
| 02 | Paywalls, Trials and Freemium | What are the three models, and which one fits your activation curve? |
| 03 | Expansion Mechanics | Seats, usage, upsell, cross-sell: which one compounds for you? |
| 04 | Dunning and Involuntary Churn | What is the cheapest revenue you are currently not collecting? |
| 05 | Monetising Without Killing Growth | How do you add a paywall without flattening the top of the funnel? |
| 06 | The Pricing Experiment | How do you test price without burning trust or breaking your data? |

### VI — Lifecycle and Automation  `[Advanced]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | The Lifecycle Map | Why model states rather than campaigns? |
| 02 | Triggered vs Scheduled | Which messages should be event-driven, and which should not exist? |
| 03 | Email, Push and In-Product | How do you pick a channel from the user's intent rather than your convenience? |
| 04 | Building a Lifecycle Program | How do you get from one email to a system that maintains itself? |
| 05 | Measuring Lifecycle | Incrementality and holdouts: is the program adding anything at all? |
| 06 | Automation Without Spam | What is a frequency budget, and how do you enforce one? |

### VII — The Growth Operating System  `[Expert]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | The Experiment Machine | Intake, prioritisation, cadence, review — what does the machine look like? |
| 02 | The Brief and the Readout | What two documents make experiments cumulative rather than forgotten? |
| 03 | Portfolio Thinking | How do you balance a few big bets against many small tests? |
| 04 | The Growth Team | How is it structured, who does it report to, and what breaks in each setup? |
| 05 | Worked: Fixing a Broken Funnel | One funnel, from diagnosis to recovered conversion, with numbers |
| 06 | Worked: Launching a Referral Loop | One loop, from k-factor model to post-launch readout |

### VIII — Growth by Sector  `[Expert]` ★
*Short deltas. Market shape lives in `product-management/XIII` — these say what it does to growth.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Enterprise B2B SaaS | Why does self-serve activation barely exist when the buyer is not the user? |
| 02 | Fintech | Which growth levers are regulated, restricted or outright illegal here? |
| 03 | Ecommerce and D2C | Why is growth a margin problem rather than a funnel problem? |
| 04 | Marketplaces | How do you grow two sides at once without breaking liquidity? |
| 05 | Devtools and Infra | Why does developer growth run on docs, trust and time-to-first-call? |
| 06 | Healthtech and Vertical SaaS | What happens when the market is small, slow and reference-driven? |

### `projects/` — Build
`01-growth-model` (Go) · `02-activation-funnel-simulator` (Go) · `03-k-factor-and-viral-loop` (Go) · `04-lifecycle-holdout-calculator` (Go)

---

# 4. `gtm-engineering` — GTM

> Taking a product to a market: sizing it, naming the buyer, positioning against the real
> alternative, picking a motion, building pipeline arithmetic that holds, and launching.

### I — Market and Buyer  `[Foundations]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | What GTM Actually Is | What problem does go-to-market solve that product does not? |
| 02 | Market Sizing: TAM, SAM, SOM | Top-down or bottom-up — which should you believe and when? |
| 03 | Segmentation | Which segments matter, versus the ones that exist only in your CRM? |
| 04 | The ICP | How do you write an ICP that actually disqualifies people? |
| 05 | The Buying Committee | Champion, economic buyer, blocker, user: who has to say yes? |
| 06 | The Trigger | Why would they buy now rather than next year? |

### II — Positioning and Messaging  `[Basic]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Positioning Is a Choice About Competition | What are you choosing when you position, and what are you giving up? |
| 02 | The Alternative You Are Really Against | Why is the competitor usually a spreadsheet, a junior hire, or nothing? |
| 03 | The Positioning Canvas | Category, segment, differentiator, proof — on one page |
| 04 | From Positioning to Messaging | How does one position become a message hierarchy without drifting? |
| 05 | The Narrative | What is the five-slide story, and why does it sell better than a feature list? |
| 06 | Naming a Category | When is category creation worth it, and when is it vanity? |

### III — Motions and Channels  `[Intermediate]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | The Four Motions | Self-serve, inside sales, enterprise, partner: what defines each? |
| 02 | Picking a Motion From ACV | What is the arithmetic that decides which motion you can afford? |
| 03 | PLG Mechanics | What does product-led actually require — and why is it not "no sales"? |
| 04 | Sales-Led Mechanics | What is a stage gate, and what does qualification actually test? |
| 05 | PLG to Enterprise | Why is this the hardest transition in software, and what has to change first? |
| 06 | Partners and Channel | Resellers, marketplaces and co-sell: when does channel add leverage? |

### IV — Pipeline and Demand  `[Intermediate]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Pipeline Arithmetic | How do you work backwards from a revenue number to this month's activity? |
| 02 | Lead Stages and Definitions | Why do MQL/SQL/SAL cause fights, and how do you define them so they don't? |
| 03 | Demand Generation vs Capture | Which one are you actually doing, and what does each cost? |
| 04 | The Handoff | Where do leads die between marketing and sales, and how do you instrument it? |
| 05 | SDR Mechanics and Economics | What does an SDR cost per meeting, and when does the model break? |
| 06 | Forecasting and the Review | What does a pipeline review look like when it is not theatre? |

### V — Launch and Lifecycle  `[Advanced]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | What a Launch Is For | What are the three tiers of launch, and which does this deserve? |
| 02 | The Launch Plan | Who owns what, by when — the checklist that prevents the usual chaos |
| 03 | Messaging and Enablement | What does sales need in order to actually sell it? |
| 04 | Pricing at Launch | How do you price something with no usage data? |
| 05 | The Post-Launch Readout | What do you measure at 7, 30 and 90 days — and what do you do about it? |
| 06 | Sunsetting and Deprecation | How do you run the launch nobody plans for? |

### VI — AI in the GTM Stack  `[Advanced]`
*AI applied to go-to-market. Building AI **products** is `ai-engineering`; this is AI as a tool of the function.*

| # | Lesson | The question it answers |
|---|---|---|
| 01 | What AI Changes About GTM | Which parts of GTM did models actually change, and which are unchanged? |
| 02 | Research, Enrichment and Scoring | Where does a model beat a rules-based account score? |
| 03 | AI SDRs and the Volume Trap | Why did mass AI outbound stop working, and what does that imply? |
| 04 | AI Answers and the Content Shift | What happens to informational traffic when answers no longer need a click? |
| 05 | AI in Calls and Coaching | What do call-intelligence tools actually improve, measurably? |
| 06 | AI in Customer Success | Health scores, QBR prep, escalation: where is the model trustworthy? |

### VII — Worked GTM Plans  `[Expert]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Zero to One Enterprise GTM | First ten enterprise customers, worked with the pipeline maths |
| 02 | A PLG Launch | From waitlist to self-serve revenue, with the funnel numbers |
| 03 | Moving Upmarket | What changes in product, pricing, motion and ops — all at once |
| 04 | Entering a Second Segment | When is it focus-breaking, and when is it the obvious next move? |
| 05 | The GTM Review | The ritual, as a checklist you can run on any plan |

### VIII — GTM by Sector  `[Expert]` ★

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Enterprise B2B SaaS | Procurement, security review, legal: what actually blocks the deal? |
| 02 | Fintech | How do compliance and risk teams enter the buying committee? |
| 03 | Ecommerce and D2C | What replaces a sales motion when there isn't one? |
| 04 | Marketplaces | How do you sell to a side that has no budget line for you? |
| 05 | Devtools and Infra | How does bottom-up adoption convert into a contract? |
| 06 | Healthtech and Vertical SaaS | Why are reference customers and associations the whole channel? |

### `projects/` — Build
`01-icp-scorecard` (Go) · `02-pipeline-model` (Go) · `03-positioning-canvas` · `04-launch-plan-template`

---

# 5. `revenue-operations` — Operations

> The machinery underneath: operating cadence, the funnel as a data model, systems of record,
> reporting you can trust, process design, and the economics of running it.

### I — Why Operations Exists  `[Foundations]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | What Breaks Without Ops | What are the five symptoms of a company that has no operations function? |
| 02 | The Three Ops Functions | RevOps, BizOps, Product Ops — what does each own and where do they overlap? |
| 03 | Systems, Processes and Data | What are the three things ops owns, and why does it own all three together? |
| 04 | Centralised vs Embedded | Which model fits which company size, and what does each fail at? |
| 05 | What an Ops Person Does All Day | What is the actual week, and what should be automated out of it? |

### II — Operating Cadence and Goals  `[Basic]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | The Operating Calendar | What happens annually, quarterly, monthly and weekly — and who runs each? |
| 02 | Goals That Work | What are OKRs for, and what are the four ways they get misused? |
| 03 | The Weekly Business Review | What is on the agenda, and what is the one rule that keeps it useful? |
| 04 | The Decision Log | Why is this the cheapest high-leverage artifact in a company? |
| 05 | Metric Ownership | Why does every number need exactly one owner, and how do you assign them? |
| 06 | Planning Without Theatre | How do you plan a quarter without producing a document nobody reads? |

### III — Funnel and Systems Architecture  `[Intermediate]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | The Funnel as a Data Model | Why are stages states rather than opinions, and what does that enforce? |
| 02 | The System of Record Question | CRM, product DB or warehouse — which one wins when they disagree? |
| 03 | Object Model Design | Account, contact, opportunity, subscription, usage: how should they relate? |
| 04 | Lifecycle Stages and Transitions | What moves a record forward, and who is allowed to move it back? |
| 05 | Territories, Routing and Assignment | How do leads get to the right person without a human deciding? |
| 06 | The Integration Map | Why do point-to-point integrations rot, and what replaces them? |

### IV — Data and Reporting  `[Intermediate]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Why Two Dashboards Disagree | What are the five causes, and how do you diagnose which one it is? |
| 02 | The Warehouse and the Semantic Layer | What problem does a semantic layer solve that a dashboard cannot? |
| 03 | Metric Definitions as Code | How do you stop a metric from having four definitions? |
| 04 | Reporting People Actually Use | What makes a report get opened twice? |
| 05 | Data Quality and the Hygiene Loop | How do you keep the CRM from decaying, as a process rather than a cleanup? |
| 06 | Self-Serve Analytics and Its Limits | When does giving everyone SQL make things worse? |

### V — Process Design and Automation  `[Advanced]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Designing a Process | Inputs, owners, SLAs, exits: what makes a process complete? |
| 02 | Writing an SOP Someone Follows | Why do most SOPs go unread, and what does a used one look like? |
| 03 | RACI and the Accountability Problem | How do you fix "everyone thought someone else had it"? |
| 04 | Automation vs Process Fix | Why does automating a bad process make it worse, faster? |
| 05 | The Tooling Decision | Build, buy or spreadsheet — and when is the spreadsheet correct? |
| 06 | Ops Debt | How do you see it, price it, and pay it down on purpose? |

### VI — Finance Operations  `[Advanced]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | The Revenue Waterfall and Who Owns It | Who produces the number, and who is allowed to restate it? |
| 02 | Billing, Invoicing and Collections | Where does revenue leak between a signed deal and cash? |
| 03 | Revenue Recognition for Operators | Enough to not be dangerous in a pricing discussion |
| 04 | The Budget and Headcount Plan | How do you turn a plan into a hiring sequence with real constraints? |
| 05 | Unit Economics in Practice | How do you compute yours from the systems you actually have? |
| 06 | The Board Pack | What goes in it, what gets asked, and what you should volunteer? |

### VII — Trust, Risk and Compliance  `[Expert]`

| # | Lesson | The question it answers |
|---|---|---|
| 01 | The Security Review as a GTM Blocker | Why does this kill more enterprise deals than price? |
| 02 | SOC 2, ISO and What They Require | What do they actually test, and what do they not prove? |
| 03 | Privacy, GDPR and Data Residency | What constrains your product, and what constrains your ops stack? |
| 04 | Vendor and Subprocessor Management | How do you onboard a vendor — especially an AI vendor — defensibly? |
| 05 | Incident Response for Non-Engineers | What is your job during an incident, and what is not? |
| 06 | Worked: Passing Enterprise Procurement | One review, end to end, with the documents and the timeline |

### VIII — Operations by Sector  `[Expert]` ★

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Enterprise B2B SaaS | What does contract complexity do to your data model and billing? |
| 02 | Fintech | Reconciliation, audit trails, regulatory reporting: what becomes non-optional? |
| 03 | Ecommerce and D2C | Why is inventory truth the hardest operational problem here? |
| 04 | Marketplaces | Payouts, disputes and tax: what does running both sides require? |
| 05 | Devtools and Infra | How do you meter and bill usage you cannot predict? |
| 06 | Healthtech and Vertical SaaS | What do compliance regimes add to every single process? |

### `projects/` — Build
`01-funnel-data-model` · `02-metric-definition-registry` · `03-sop-and-raci-pack` · `04-board-pack-template`

---

# 6. `ai-engineering` — AI

> Building products with a model inside them. **Hard scope boundary: the product surface, not the
> infrastructure under it** — no serving, GPUs, training or cluster design.
>
> Full lesson tree is in [`PLAN.md` §2b](PLAN.md). Summary of the ladder:

| Level | Topic | What it is |
|---|---|---|
| Foundations | `I-what-a-model-changes` | The five facts about a model; why an AI feature is not a normal feature; when not to use one |
| Basic | `II-model-primitives` | Tokens and context, latency users feel, cost per request, structured output, choosing a model, non-determinism |
| Intermediate | `III-context-and-knowledge` | The prompt as a spec, context engineering, retrieval and grounding, citations, finetuning (rarely), staleness |
| Intermediate | `IV-workflows-and-agents` | The autonomy spectrum, when a workflow beats an agent, tools, approval gates, blast radius, multi-step reliability |
| Advanced | `V-evaluation-and-quality` | **Flagship.** Datasets, graders, the statistics, the release gate, reading traces |
| Advanced | `VI-ai-ux-and-trust` | The trust budget, showing uncertainty, designing for wrong answers, latency as a design material, consent, feedback loops |
| Expert | `VII-shipping-and-economics` | Cost per task vs price, pricing AI, margin under load, canary, monitoring, safety and policy in-product |
| Expert | `VIII-worked-ai-features` | A copilot, a support assistant, an analyst feature, the AI feature review |
| Expert | `IX-ai-by-sector` ★ | Where a model may and may not decide, per sector — and what that permission costs |

### `projects/` — Build
`01-eval-harness` (Go) · `02-cost-model` (Go) · `03-retrieval-quality-bench` (Go) · `04-trust-budget-calculator` (Go)

---

# 7. Sectors — how they're handled now

The standalone `sector-playbooks` course is **dissolved**. Sector content now lives inside the
course whose mechanism it modifies, so a reader finds it where they are already reading.

**`product-management/XIII-product-by-sector` is the primary home** — seven full lessons, and its
`README.md` holds the cross-sector comparison matrix that everything else links to.

Each other course gets one sector topic at Expert level, written as **short deltas** (600–1,200
words per sector) that link back to the product matrix for market shape:

| Course | Sector topic | Shape |
|---|---|---|
| `product-management` | `XIII-product-by-sector` | 7 **full** lessons + the matrix in its README |
| `growth-engineering` | `VIII-growth-by-sector` | 1 delta lesson per sector: which growth levers are unavailable, illegal or inverted |
| `gtm-engineering` | `VIII-gtm-by-sector` | 1 per sector: buyer, motion, cycle, procurement, what blocks a deal |
| `revenue-operations` | `VIII-operations-by-sector` | 1 per sector: the processes, systems and controls the sector forces on you |
| `ai-engineering` | `IX-ai-by-sector` | 1 per sector: where a model may and may not decide, and what that costs |
| `metrics-and-models` | `VII-benchmarks-and-their-limits/02` | the benchmark table, sourced and dated |

**Why this is better than a separate course:** a sector-only course has no spine of its own — it
is entirely deltas, which makes it unreadable on its own and a magnet for duplication. Attaching
deltas to the mechanism they modify also means a sector can be added incrementally: write the
product lesson first, add the growth and GTM deltas when you need them.

## Sector roadmap

**Tier 1 — ship these** (the seven in `XIII`, chosen for structural contrast):

Enterprise B2B SaaS · Fintech · Ecommerce & D2C · Marketplaces · Devtools & Infra ·
Healthtech & Regulated Care · Vertical SaaS

**Tier 2 — strong candidates for later:**

| Sector | Why it's worth a lesson |
|---|---|
| Consumer subscription & media | Content economics, churn seasonality, the free tier as the whole funnel |
| Logistics & supply chain | Physical-world constraints, exception handling as the actual product |
| Insurtech | Underwriting as product, regulated pricing, the claims moment |
| Proptech | Long cycles, offline incumbents, and transactions that happen once |
| Edtech | Buyer ≠ user ≠ payer, institutional cycles, seasonality |
| Govtech & public sector | Procurement as the entire game, accessibility mandates, multi-year cycles |
| Climate & energy | Hardware coupling, policy dependence, measurement as product |
| Gaming | Economy design, live-ops cadence, retention mechanics unlike anything else |
| Hardware + software | Inventory risk, firmware release cadence, no hotfixes |

**Note on scope:** "AI-native products" is deliberately **not** a sector. It's handled by the
`ai-engineering` course and the `## The AI layer` section in every lesson. Treating it as a
sector would make it a bucket for everything.

---

## Cross-course rules (enforced in review)

1. **Metrics** are defined only in `metrics-and-models`. Every other lesson links.
2. **Sector market shape** is defined only in `product-management/XIII-product-by-sector`
   (its `README.md` holds the matrix). Other courses' sector topics contain *function deltas only*
   and link back to it — they never restate buyer, cycle length or regulation.
   **Archetype comes before sector:** `product-management/IX-product-archetypes` is a prerequisite
   for every sector lesson, because shape usually dominates market.
3. **AI products** are covered only in `ai-engineering`. AI *applied to a function* lives in
   that function's `## The AI layer` section and in `gtm-engineering/VI-ai-in-the-gtm-stack`.
4. **Pricing** is split deliberately: mechanism in `product-management/V-pricing-and-packaging`,
   the market-facing decision in `gtm-engineering`, the arithmetic in `metrics-and-models`.
   Each links to the other two rather than repeating them.
5. Every course ends with a **How it connects to the other courses** section in its `README.md`,
   the same convention the reference repo uses.
