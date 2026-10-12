# Glossary

The terms this course uses, defined the way it uses them. Where a word is contested in the wider
industry, the entry says so rather than pretending there is a settled meaning. The right hand
column names the topic that treats the term properly.

## The job

| Term | What it means here | Topic |
|---|---|---|
| **Closure rate** | The share of shipped work that received a written readout against a bar set beforehand. Usually between 10% and 30% the first time a team measures it | The Product Function |
| **Cycle time** | One full turn of the product loop, from "we believed this" to "we know whether that was right". Not engineering throughput | The Product Function |
| **Decision memo** | One page recording the options, the call, the reasoning and what would prove it wrong. The cheapest artifact and the most skipped | The Product Function |
| **Falsification test** | The thing that, if observed, would mean the decision was wrong. A decision without one cannot be learned from | The Product Function |
| **Refusal** | Work deliberately not done, recorded with the reason. A quarter with no refusals is a backlog being executed, not a roadmap | The Product Function |
| **Behaviour contract** | The replacement for a specification when output varies: must, must never, should, may | Specs for Non Deterministic Systems |

## What a model is

| Term | What it means here | Topic |
|---|---|---|
| **Context window** | The total text a model can hold at once, counted in tokens. A budget, not a setting | Model Primitives |
| **Token** | The unit models read, write and charge by. Roughly three quarters of a word in English | Model Primitives |
| **Time to first token** | Seconds before any output appears. The latency users actually feel on a streaming surface | Model Primitives |
| **Temperature** | The knob controlling how much randomness is allowed in sampling. Zero is not deterministic in practice | Model Primitives |
| **Structured output** | Output constrained to a schema so a program can rely on it, through tool calls or a JSON mode | Model Primitives |
| **Grounding** | Supplying the facts an answer must rest on, so the model composes rather than recalls | Context and Knowledge |
| **Retrieval** | Fetching relevant material into the context window at request time. Often called RAG | Context and Knowledge |
| **Context engineering** | Deciding what goes into the window, in what order, at what cost. Mostly a product decision | Context and Knowledge |
| **Prompt injection** | Text encountered while working that the model treats as instruction. There is no parser boundary to defend | Trust, Safety and Guardrails |

## Product shapes and autonomy

| Term | What it means here | Topic |
|---|---|---|
| **Product shape** | Which of the six forms a product takes, decided by who reads the output before anything happens because of it | What a Model Changes |
| **Autonomy spectrum** | Suggest, draft, act with approval, act. Where a product sits determines the accuracy it needs | Agents and Autonomy |
| **Agent** | A loop where a model chooses actions, observes results and continues. Distinct from a workflow, where the steps are fixed | Agents and Autonomy |
| **Workflow** | A fixed sequence of steps, some of which call a model. Cheaper, more reliable, and usually the right answer | Agents and Autonomy |
| **Tool** | Something a model may call to act or to look something up. The set of tools is the set of things it can do wrong | Agents and Autonomy |
| **Blast radius** | What a wrong action can reach before anyone notices. The thing you design, not the accuracy | Agents and Autonomy |
| **Approval fatigue** | The decay of a human review step into rubber stamping. A rejection rate near zero is the alarm | What a Model Changes |
| **Trust budget** | The number of wrong answers a user will absorb before deciding the product is a thing they check rather than trust | AI Product UX |

## Evaluation

| Term | What it means here | Topic |
|---|---|---|
| **Eval set** | The cases a change is measured against. The real specification of an AI product | Evaluation and Quality |
| **Grader** | What decides whether one output passed: code, a model judge, or a human | Evaluation and Quality |
| **Judge** | A model scoring another model's output against a rubric. Worthless until validated against human labels | Evaluation and Quality |
| **Pass rate** | The share of cases that passed. Needs a confidence interval and repeated runs to mean anything | Evaluation and Quality |
| **Release gate** | The bar a change must clear before shipping: quality not worse, must pass cases passing, cost and latency inside budget | Evaluation and Quality |
| **Trace** | The stored record of one request: prompt after templating, retrieved documents, model version, tool calls, output, scores | AI Product Org and Process |
| **Groundedness** | Whether every claim in an answer is supported by the sources supplied. Measured separately from whether it is helpful | AI Quality Metrics |
| **Drift** | Quality moving without anyone changing anything, usually because the model behind it changed | Model Lifecycle and Drift |

## Money and measurement

| Term | What it means here | Topic |
|---|---|---|
| **CAC** | Fully loaded acquisition spend divided by customers won. Paid against paid, total against total, never one of each | Revenue and Efficiency |
| **Payback** | Months of gross profit needed to repay acquisition cost. Two honest versions, differing by a third or more | Revenue and Efficiency |
| **LTV** | A forecast, not a measurement. Depends entirely on which margin, which churn and what horizon | Revenue and Efficiency |
| **Gross margin** | Revenue minus cost of revenue, which includes inference. Stops being a constant once a model is in the product | AI Cost and Margin |
| **Cost per task** | What one unit of useful work costs in inference, retries, judging and human review | AI Cost and Margin |
| **Value metric** | The unit you charge for. The single highest leverage pricing decision | Pricing and Packaging |
| **Containment** | The share of requests resolved without a human. The support metric most often quoted and least often qualified | AI Quality Metrics |
| **Acceptance rate** | The share of suggestions a user keeps. The primary metric of the suggestion shape | AI Quality Metrics |
| **Intervention rate** | How often a human has to step in. Rising intervention is the first sign of a quality problem in an agent | Agent and Task Metrics |
| **Cohort** | A group defined by when they arrived, followed over time. The only honest way to read retention | Funnel, Retention and Experiments |
| **North star** | One metric a team orients around. Useful for focus, misleading whenever the business has two sides | How to Read a Metric |

## Selling and sectors

| Term | What it means here | Topic |
|---|---|---|
| **ICP** | The profile of a customer worth having. Written properly, it disqualifies | Product by Sector |
| **Buying committee** | Everyone who must agree: champion, economic buyer, procurement, security, legal. In enterprise, rarely the user | Product by Sector |
| **Admin surface** | Single sign on, provisioning, roles, audit logs, residency and retention. Nobody asks for it in a demo and every security review requires it | Product by Sector |
| **Security review** | A product backlog arriving through the sales channel. Keep a register with deal value attached | Product by Sector |
| **Subprocessor** | A third party that touches customer data on your behalf, including a model provider. Disclosed, and increasingly contractual | Product by Sector |
| **Take rate** | The share of transaction value a marketplace keeps. A product lever, not a finance setting | Product by Sector |
